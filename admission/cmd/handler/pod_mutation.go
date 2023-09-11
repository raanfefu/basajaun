package handler

import (
	"log"
	"os"

	"github.com/raanfefu/basajaun/admission/cmd/k8s"
	"github.com/raanfefu/basajaun/admission/pkg"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const NAME_ANNOTATIONS_BY_SECRET_AZ_CONFIG = "opa-cfg-file"
const NAME_ANNOTATIONS_BY_BUNDLE_NAME = "opa-bundle-name"
const NAME_CONFIGMAP_BY_FILE_CONFIG_ENVOY = "proxy-config"
const ENVOY_CFG_PATH = "/envoy-cfg"
const OPA_CFG_PATH = "/app/env/"
const OPA_CONTAINER_NAME = "opa"

func PodMutationHandler(admissionReview *admissionv1.AdmissionReview) *admissionv1.AdmissionReview {
	annotations := k8s.DecodeAnnotation(admissionReview)
	admissionReviewOut := preparedAdmisionResponse(admissionReview)
	_, ok := annotations["opa-enabled"]
	if !ok {
		log.Printf("Skip authz inject must be annotations opa-enabled = true")
		admissionReviewOut.Response.Allowed = true
		admissionReviewOut.Response.Patch = k8s.EncodeBase64Patches(make([]pkg.Patch, 0))
		return admissionReviewOut
	}

	bundle_name_value, ok := annotations[NAME_ANNOTATIONS_BY_BUNDLE_NAME]
	if !ok {
		log.Printf("reject to set annotations %s \n", NAME_ANNOTATIONS_BY_BUNDLE_NAME)
		admissionReviewOut.Response.Allowed = false
		return admissionReviewOut
	}

	secret_value_az_config, ok := annotations[NAME_ANNOTATIONS_BY_SECRET_AZ_CONFIG]
	if !ok {
		log.Printf("reject to set annotations %s \n", NAME_ANNOTATIONS_BY_SECRET_AZ_CONFIG)
		admissionReviewOut.Response.Allowed = false
		return admissionReviewOut
	}

	_, err := k8s.GetSecret(admissionReview.Request.Namespace, secret_value_az_config)
	if err != nil {
		log.Printf("err get secret  %s\n ", secret_value_az_config)
		admissionReviewOut.Response.Allowed = false
		return admissionReviewOut
	}

	_, err = k8s.GetConfigmap(admissionReview.Request.Namespace, NAME_CONFIGMAP_BY_FILE_CONFIG_ENVOY)
	if err != nil {
		log.Printf("err get configmap  %s\n ", NAME_CONFIGMAP_BY_FILE_CONFIG_ENVOY)
		admissionReviewOut.Response.Allowed = false
		return admissionReviewOut
	}

	patches := make([]pkg.Patch, 0)

	/* agregar al patches los init container */
	init := initContainers()
	patches = k8s.AddPatches(patches, "add", "/spec/initContainers", init)

	/* agrear al patches el contenedor de envoy */
	volumeEnvoyConfig := k8s.VolumeConfigMap(NAME_CONFIGMAP_BY_FILE_CONFIG_ENVOY)
	proxy := containerEnvoy(admissionReview, nil)

	patches = k8s.AddPatches(patches, "add", "/spec/volumes/-", volumeEnvoyConfig)
	patches = k8s.AddPatches(patches, "add", "/spec/containers/-", proxy)

	/* agregar al patches los opa container */
	volumeOpaConfig := k8s.VolumeSecret(secret_value_az_config)
	patches = k8s.AddPatches(patches, "add", "/spec/volumes/-", volumeOpaConfig)

	enviroment := make([]corev1.EnvVar, 0)
	enviroment = append(enviroment,
		corev1.EnvVar{
			Name:  "BUNDLE_NAME",
			Value: bundle_name_value,
		})

	rules := containerOpa(secret_value_az_config, &enviroment)
	patches = k8s.AddPatches(patches, "add", "/spec/containers/-", rules)

	/* completar el response para validar el mutating */
	admissionReviewOut.Response.Allowed = true
	admissionReviewOut.Response.Patch = k8s.EncodeBase64Patches(patches)

	return admissionReviewOut
}

func preparedAdmisionResponse(admissionReview *admissionv1.AdmissionReview) *admissionv1.AdmissionReview {
	patchType := admissionv1.PatchTypeJSONPatch
	admissionResponse := &admissionv1.AdmissionResponse{
		UID:       admissionReview.Request.UID,
		PatchType: &patchType,
	}
	admissionReview.Response = admissionResponse
	return admissionReview
}

func containerEnvoy(admissionReview *admissionv1.AdmissionReview, env *[]corev1.EnvVar) *corev1.Container {

	volumeMounts := append(make([]corev1.VolumeMount, 0),
		corev1.VolumeMount{
			Name:      NAME_CONFIGMAP_BY_FILE_CONFIG_ENVOY,
			MountPath: ENVOY_CFG_PATH,
			ReadOnly:  true,
		})
	runAsUser := int64(111)

	sidecar := &corev1.Container{
		Name:         "envoy",
		Image:        os.Getenv("PROXY_IMAGE"),
		Args:         []string{"envoy", "--config-path", "/envoy-cfg/envoy.yaml"},
		VolumeMounts: volumeMounts,
		SecurityContext: &corev1.SecurityContext{
			RunAsUser: &runAsUser,
		},
	}
	if env != nil {
		sidecar.Env = *env
	}

	return sidecar
}

func containerOpa(volumeName string, env *[]corev1.EnvVar) *corev1.Container {

	volumeMounts := make([]corev1.VolumeMount, 0)
	volumeMounts = append(make([]corev1.VolumeMount, 0), corev1.VolumeMount{
		Name:      volumeName,
		MountPath: OPA_CFG_PATH,
		ReadOnly:  true,
	})

	runAsUser := int64(111)
	sidecar := &corev1.Container{
		Name:  "opa",
		Image: os.Getenv("OPA_IMAGE"),
		//	Env:             *env,
		ImagePullPolicy: corev1.PullAlways,
		SecurityContext: &corev1.SecurityContext{
			RunAsUser: &runAsUser,
		},
		VolumeMounts: volumeMounts,
		LivenessProbe: &corev1.Probe{
			InitialDelaySeconds: 5,
			PeriodSeconds:       5,
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Port: intstr.IntOrString{
						IntVal: 8282,
					},
					Scheme: "HTTP",
					Path:   "/health?plugins",
				},
			},
		},
		ReadinessProbe: &corev1.Probe{
			InitialDelaySeconds: 5,
			PeriodSeconds:       5,
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Port: intstr.IntOrString{
						IntVal: 8282,
					},
					Scheme: "HTTP",
					Path:   "/health?plugins",
				},
			},
		},
	}
	if env != nil {
		sidecar.Env = *env
	}

	return sidecar
}

func initContainers() []corev1.Container {
	False := false
	rootUser := int64(0)
	initContainer := &corev1.Container{
		Name:  "init-container-opa",
		Image: os.Getenv("PROXY_INIT"),
		Args:  []string{"-p", "8000", "-u", "1111", "-w", "8282"},
		SecurityContext: &corev1.SecurityContext{
			Capabilities: &corev1.Capabilities{
				Add: []corev1.Capability{
					"NET_ADMIN",
				},
			},
			RunAsNonRoot: &False,
			RunAsUser:    &rootUser,
		},
	}

	r := make([]corev1.Container, 1)
	r[0] = *initContainer
	return r
}

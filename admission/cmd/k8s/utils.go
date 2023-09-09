package k8s

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetSecret(namespace, name string) (*map[string][]byte, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Failed to create K8s config client")
		return nil, err

	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to create K8s clientset")
		return nil, err

	}
	secret, err := client.CoreV1().Secrets(namespace).Get(context.TODO(), name, meta.GetOptions{})
	if err != nil {
		log.Printf("Failed to get K8s Secret %s", err)
		return nil, err
	}
	return &secret.Data, nil
}

func GetConfigmap(namespace string, name string) (*map[string][]byte, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Failed to create K8s config client")
		return nil, err

	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to create K8s clientset")
		return nil, err

	}
	configmap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, meta.GetOptions{})
	if err != nil {
		log.Printf("Failed to get K8s Secret %s", err)
		return nil, err
	}
	return &configmap.BinaryData, nil
}

func VolumeSecret(secretName string) *corev1.Volume {
	var volume *corev1.Volume

	volume = &corev1.Volume{
		Name: secretName,
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName: secretName,
			},
		},
	}
	return volume
}

func VolumeConfigMap(configMapName string) *corev1.Volume {

	volume := &corev1.Volume{
		Name: configMapName,
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: configMapName,
				},
			},
		},
	}
	return volume
}

func DecodeAnnotation(admissionReview *admissionv1.AdmissionReview) map[string]string {
	if admissionReview != nil {
		data := &corev1.Pod{}
		json.Unmarshal(admissionReview.Request.Object.Raw, data)
		return data.Annotations
	} else {
		return nil
	}
}

func DecodeAdmissionReview(r *http.Request) admissionv1.AdmissionReview {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := &admissionv1.AdmissionReview{}

	json.Unmarshal(b, data)
	return *data
}

#!/bin/sh


REGISTRY=`cat  ../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.registry`
ADMISSION_CRTL_IMAGE=`cat  ../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.image_admission_ctrl`

docker push ${REGISTRY}/${ADMISSION_CRTL_IMAGE}
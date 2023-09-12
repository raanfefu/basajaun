#!/bin/sh


REGISTRY=`cat  ../charts/sample/values.yaml | yq | jq -r .json_values.deployment.registry`
ADMISSION_CRTL_IMAGE=`cat  ../charts/sample/values.yaml | yq | jq -r .json_values.deployment.image_sample`

docker push ${REGISTRY}/${ADMISSION_CRTL_IMAGE}
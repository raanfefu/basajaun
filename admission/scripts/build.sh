#!/bin/sh
 
GOOS=`cat  ../charts/admission-controller/values.yaml | yq | jq -r .json_values.build_values.golang.goos`
GOARCH=`cat  ../charts/admission-controller/values.yaml | yq | jq -r .json_values.build_values.golang.goarch`
APP_NAME=`cat  ../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.name`
REGISTRY=`cat  ../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.registry`
ADMISSION_CRTL_IMAGE=`cat  ../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.image_admission_ctrl`
DOCKER_PLATFORM=${GOOS}/${GOARCH}

if test -d ${BUILD_PATH}; then
    rm -rf ${BUILD_PATH}
fi

mkdir ${BUILD_PATH}
if [ $? -eq 0 ]; then
    echo "Starting build golang binary file: \"${APP_NAME}\" ..."
    env GOOS=${GOOS} GOARCH=${GOARCH}  go build -o ${BUILD_PATH}/${APP_NAME} ./cmd/main.go 
else 
    exit -1
fi

if [ $? -eq 0 ]; then
    echo "Finish build golang binary file: \"${APP_NAME}\" ...\n"
else 
    exit -1
fi

cp ./docker/* ${BUILD_PATH}
if [ $? -eq 0 ]; then
    echo "Starting build docker image local: \"${REGISTRY}/${ADMISSION_CRTL_IMAGE}\" by platform ${DOCKER_PLATFORM} ..."
    cd ${BUILD_PATH} && docker build --platform=${DOCKER_PLATFORM} -t ${REGISTRY}/${ADMISSION_CRTL_IMAGE} -f Dockerfile.controller .
else
    exit -1
fi
 
if [ $? -eq 0 ]; then
    echo "Finish build image docker local: \"${REGISTRY}/${ADMISSION_CRTL_IMAGE}\" by platform ${DOCKER_PLATFORM} ..."
    
else 
    exit -1
fi

if test -d ${BUILD_PATH}; then
    rm -rf ${BUILD_PATH}
fi
 


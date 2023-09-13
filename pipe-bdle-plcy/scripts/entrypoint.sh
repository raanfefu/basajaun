#!/bin/bash


build_custom () {
    echo "Rest Policy  !!! "
    rm -rf dist && mkdir dist
    opa build src --output dist/`cat manifest.json | jq -r ".name"`.tar.gz
    upload_artifact 
}

build_rest (){
    echo "Rest Policy  !!! "
    rm -f src/policy.rego
    rm -rf dist && mkdir dist
    check-jsonschema --schemafile /jsonschema/schema-rest.json rest.json
    jinja2 /template/rest-policy-template.jinja2 rest.json --format=json 
    opa build src --output dist/`cat manifest.json | jq -r ".name"`.tar.gz 
    upload_artifact 
}

upload_artifact(){
    export AZCOPY_AUTO_LOGIN_TYPE=SPN
    export AZCOPY_SPA_APPLICATION_ID="1389cf37-c824-446c-802a-096e1fbd535a"
    export AZCOPY_SPA_CLIENT_SECRET='cCy8Q~pr~NvM~DWhXsBXk5p24oM4i69hPIwqRcKy'
    export AZCOPY_TENANT_ID="27eb7bd2-49a9-40dd-bc1d-516ab57c8258"
    azcopy rm https://opabundletest.blob.core.windows.net/bundle/`cat manifest.json | jq -r ".name"`.tar.gz
    azcopy cp ./dist/`cat manifest.json | jq -r ".name"`.tar.gz https://opabundletest.blob.core.windows.net/bundle/
}

build () {
    TYPE=`cat manifest.json | jq -r .type`
     
    if [[ "REST" = "${TYPE}" ]]; then
        build_rest
    fi

     if [[ "CUSTOM" = "${TYPE}" ]]; then
        build_custom
    fi
} 
 
if [ -f manifest.json ]; then
    check-jsonschema --schemafile /jsonschema/schema-manifest.json manifest.json
    if [ $? -eq 0 ]; then
        build
    else
        echo "Error manifest.json"
    fi
else
    echo "NOT FOUND manifest.json" 
fi


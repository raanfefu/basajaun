#!/bin/sh

 
echo "Starting building angular"
cd  angular/ && ng build 

if [ $? -eq 0 ]; then
    echo "Finish build angular"
else
    exit -1
fi
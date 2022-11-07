#!/bin/bash

mkdir -p ./spec

python3 -m grpc_tools.protoc \
  --python_out=./spec \
  --grpc_python_out=./spec \
  -I ../proto/ \
  ../proto/*.proto

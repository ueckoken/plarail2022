#!/bin/bash

mkdir -p ./spec

npx grpc_tools_node_protoc \
  --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
  --js_out=import_style=commonjs,binary:./spec \
  --grpc_out=grpc_js:./spec \
  --ts_out=service=grpc-node,mode=grpc-js:./spec \
  -I ../proto/ \
  ../proto/*.proto


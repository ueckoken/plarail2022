#!/bin/bash

mkdir -p ./spec

python3 -m grpc_tools.protoc \
  --python_out=./spec \
  --grpc_python_out=./spec \
  -I ../../proto/ \
  ../../proto/*.proto

# change import to from . import on block_pb2_grpc.py
sed -i -e 's/import block_pb2/from . import block_pb2/g' ./spec/block_pb2_grpc.py
rm ./spec/block_pb2_grpc.py-e
sed -i -e 's/import ats_pb2/from . import ats_pb2/g' ./spec/ats_pb2_grpc.py
rm ./spec/ats_pb2_grpc.py-e
sed -i -e 's/import statesync_pb2/from . import statesync_pb2/g' ./spec/statesync_pb2_grpc.py
rm ./spec/statesync_pb2_grpc.py-e
sed -i -e 's/import speedControl_pb2/from . import speedControl_pb2/g' ./spec/speedControl_pb2_grpc.py
rm ./spec/speedControl_pb2_grpc.py-e
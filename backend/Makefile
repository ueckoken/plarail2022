up:
	docker compose down
	docker compose up --build -d 

pull:
	git pull

test:
	docker compose -f compose.debug.yml up --build
start: pull up

pb: statesync speedControl ats block
	
statesync: statesync_internal statesync_external statesync_positioning

statesync_internal:
	protoc \
	--go_out=internal \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=internal/spec \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto

statesync_external:
	protoc \
	--go_out=external \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=external/spec \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto

statesync_positioning:
	protoc \
	--go_out=positioning \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=positioning/spec \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto

speedControl: speedControlSpeed speedControlPositioning

speedControlSpeed:
	protoc \
	--go_out=speed/spec \
	--go_opt=Mproto/speedControl.proto=. \
	--go-grpc_out=speed/spec \
	--go-grpc_opt=Mproto/speedControl.proto=. \
	-I./ \
	proto/speedControl.proto

speedControlPositioning:
	protoc \
	--go_out=positioning/spec \
	--go_opt=Mproto/speedControl.proto=. \
	--go-grpc_out=positioning/spec \
	--go-grpc_opt=Mproto/speedControl.proto=. \
	-I./ \
	proto/speedControl.proto

ats: ats_node ats_go ats_python

ats_node:
	cd ats && \
	./protoc-gen.sh

ats_go:
	protoc \
	--go_out=json2grpc/spec \
	--go_opt=Mproto/ats.proto=. \
	--go-grpc_out=json2grpc/spec \
	--go-grpc_opt=Mproto/ats.proto=. \
	-I./ \
	proto/ats.proto

ats_python:
	cd auto_operation && \
	./protoc-gen.sh

block: block_external

block_external:
	protoc \
	--go_out=external \
	--go_opt=Mproto/block.proto=./spec \
	--go-grpc_out=external/spec \
	--go-grpc_opt=Mproto/block.proto=. \
	-I./ \
	proto/block.proto

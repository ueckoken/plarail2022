up:
	docker-compose down
	docker-compose build
	docker-compose up -d

pull:
	git pull

start: pull up

pb: statesync speedControl
	
statesync: statesync_internal statesync_external statesync_positioning

statesync_internal:
	protoc \
	--go_out=internal \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=internal \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto

statesync_external:
	protoc \
	--go_out=external \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=external \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto

statesync_positioning:
	protoc \
	--go_out=positioning \
	--go_opt=Mproto/statestnc.proto=. \
	--go-grpc_out=positioning \
	--go-grpc_opt=Mproto/statesync.proto=. \
	-I./ \
	proto/statesync.proto

speedControl: speedControl_builtin speedControlSpeed speedControlPositioning

speedControl_builtin:
	protoc \
	--go_out=builtin/speed/spec \
	--go_opt=Mproto/speedControl.proto=. \
	--go-grpc_out=builtin/speed/spec \
	--go-grpc_opt=Mproto/speedControl.proto=. \
	-I./ \
	proto/speedControl.proto

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

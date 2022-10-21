up:
	docker-compose down
	docker-compose build
	docker-compose up -d

pull:
	git pull

start: pull up

pb: statesync speedControl ats
	
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

ats: ats-node
ats-node:
	cd ats && \
	./protoc-gen.sh

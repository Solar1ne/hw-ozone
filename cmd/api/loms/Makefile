LOCAL_BIN=$(CURDIR)\bin
PROTO_PATH=cmd\api\loms\v1

PROTOC=PATH="$$PATH;$(LOCAL_BIN)" protoc

protoc-generate:
	$(PROTOC) ^
	-I $(PROTO_PATH) ^
	-I vendor-proto ^
	--plugin=protoc-gen-go=$(LOCAL_BIN)\protoc-gen-go.exe ^
	--go_out=pkg\$(PROTO_PATH) ^
	--go_opt=paths=source_relative ^
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)\protoc-gen-go-grpc.exe ^
	--go-grpc_out=pkg\$(PROTO_PATH) ^
	--go-grpc_opt=paths=source_relative ^
	--plugin=protoc-gen-validate=$(LOCAL_BIN)\protoc-gen-validate.exe ^
	--validate_out="lang=go,paths=source_relative:pkg\api\loms\v1" ^
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)\protoc-gen-grpc-gateway.exe ^
	--grpc-gateway_out=pkg\$(PROTO_PATH) ^
	--grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true ^
	$(PROTO_PATH)\loms.proto
	go mod tidy

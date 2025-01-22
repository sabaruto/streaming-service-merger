all: gen-auth-proto	

gen-auth-proto:
	@protoc \
        --proto_path=protobuf \
        --go_out=services/common/genproto/authorisation \
        --go_opt=paths=source_relative \
        --go-grpc_out=services/common/genproto/authorisation \
        --go-grpc_opt=paths=source_relative \
        "protobuf/authorisation.proto"

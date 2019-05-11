.PHONY: generate

generate:
	protoc --proto_path=proto --go_out=plugins=grpc:proto pds.proto
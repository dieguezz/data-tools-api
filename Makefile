create_proto: 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/nif.proto
	protoc -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative proto/nif.proto
	protoc -I . --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true proto/nif.proto
	protoc -I . --openapiv2_out . --openapiv2_opt logtostderr=true proto/nif.proto
clean_proto:
	rm proto/*.go

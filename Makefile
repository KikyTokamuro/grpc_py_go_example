PROTOC = protoc 
PYTHON = python3

PROTO_DIR = dir_watcher

.DEFAULT_GOAL := default

default: build_go_protobuf build_py_protobuf build_server

clean:
	rm $(PROTO_DIR)/*.go
	rm $(PROTO_DIR)/*.py
	rm server/server
	rm -rf $(PROTO_DIR)/__pycache__/

build_go_protobuf:
	$(PROTOC) --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

build_py_protobuf:
	$(PYTHON) -m grpc_tools.protoc -I=. --python_out=. \
		--grpc_python_out=. $(PROTO_DIR)/*.proto

build_server:
	go build -o server/server server/server.go


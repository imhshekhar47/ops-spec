dependency:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

clean:
	if [ -d "./pb" ]; then rm -rf ./pb/*pb.go; rm -rf ./pb/*gw.go; fi;
	if [ -d "./build/openapi/v2" ]; then rm -rf "./build/openapi/v2"; fi;
	if [ -d "./build/tests" ]; then rm -rf "./build/tests"; fi;


generate: clean
	if [ ! -d "./pb" ]; then mkdir -p "./pb"; fi;
	if [ ! -d "./build/openapi/v2" ]; then mkdir -p "./build/openapi/v2"; fi;

	protoc -I="proto" \
        --go_out="pb" --go_opt=paths=source_relative \
        --go-grpc_out="pb" --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out="pb" --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out ./build/openapi/v2 \
        --openapiv2_opt logtostderr=true \
        proto/*.proto

test: generate
	if [ ! -d "./build/tests" ]; then mkdir -p "./build/tests"; fi;
	gotestsum --format testname \
		--junitfile "./build/tests/unit-tests.xml" \
		-- -coverprofile=build/coverage.out ./...

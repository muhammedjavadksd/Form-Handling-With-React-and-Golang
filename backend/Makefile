grpc: grpc-form grpc-api


grpc-form:
	protoc form-manager/pkg/server/pb/form.proto --go_out=./ --go-grpc_out=./ 

grpc-api:
	protoc api-gateway/pkg/form-client/form/pb/form.proto --go_out=. --go-grpc_out=.																


run:
	cd api-gateway && gnome-terminal -- go run main.go &
	cd form-manager && gnome-terminal -- go run main.go &

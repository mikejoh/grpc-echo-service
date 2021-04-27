
gen-cert:
	cd cert/; ./gen-cert.sh; cd ..

gen-proto:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative echo/echo.proto

server:
	go run echo_server/*.go

client:
	go run echo_client/*.go

kind-cluster:
	kind create cluster --name k8s-1.19.7 --config kind/k8s-1.19.7.yaml

install-nginx-ingress:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml

docker-build:
	docker build . -t mikejoh/grpc-echo-service:v0.1.0

docker-push:
	docker push mikejoh/grpc-echo-service:v0.1.0


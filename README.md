# gRPC echo service

_Heavily inspired by the official grpc [`helloworld`](https://grpc.io/docs/languages/go/quickstart/) example but with a fully working Kubernetes deployment example._

_Will only work on your machine._

The echo-service works as follows: The server will answer `pong` on any `ping` request sent using the gRPC client. The `pong` response and `ping` message are hardcoded in the server and client.

## Prerequisites
* Docker
* [`kind`](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
* `protoc` and Go plugins for the protocol compiler
* Go
## How to
1. See the prerequisites [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites) to install `protoc` and the Go plugins needed.
2. In the root of this repository run:
```
make gen-proto
```
3. Start the `kind` cluster:
```
kind create cluster --name k8s-1.19.7 --config kind/k8s-1.19.7.yaml
```
4. Install `metallb` to be able to create a `Service` of type `LoadBalancer` in Kubernetes:
```
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/namespace.yaml
kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)" 
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/metallb.yaml
```
5. Identify the CIDR network used by Docker:
```
docker network inspect -f '{{.IPAM.Config}}' kind
```
6. Create and apply a `ConfigMap` using a subset of the addresses in the CIDR from the output above, you should replace the IP addresses `addresses` field below:
```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: metallb-system
  name: config
data:
  config: |
    address-pools:
    - name: default
      protocol: layer2
      addresses:
      - 172.18.255.200-172.18.255.250
EOF
```
7. Build the grpc-echo-service Docker image:
```
docker build . -t grpc-echo-service:v0.1.0
```
8. Push the built grpc-echo-service Docker image to a registry.
9. Change the image in `manifests/deployment.yaml` pointing to the one you've just pushed.
10. Apply the `deployment.yaml` and `service.yaml` files.
```
kubectl apply -f manifests/deployment.yaml
kubectl apply -f manifests/service.yaml
```
11. Take a note of the Load Balancer IP the `Service` have received from MetalLB:
```
kubectl get svc grpc-echo-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}'
```
12. Add the Load Balancer IP to the `cert/echo_client-ext.cnf` and `cert/echo_server-ext.cnf` files.
13. Re-generate the certificates:
```
make gen-cert
```
14. Re-build and re-push the Docker image.
15. Restart the Deployment of grpc-echo-service server component:
```
kubectl deployment rollout restart grpc-echo-service
```
16. Run the gRPC client:
```
go run echo_client/*.go -addr "172.18.255.200:8443" -cert cert/client-cert.pem -key cert/client-key.pem
```
you should see an output similar to this one:
```
2021/05/02 22:40:35 Starting gRPC echo service client..
2021/05/02 22:40:35 Echoes service replied: pong
```

# Saga POC

The POC demonstartes the approach to build micro service oriented arhitecture on the bases of the [gRPC](https://grpc.io/docs/languages/go/basics/) framework designed by Google. 

## Tools
### [gRPC](https://grpc.io/docs/languages/go/basics/)
### [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
### [buf](https://github.com/bufbuild/buf)
### [Skaffold](https://skaffold.dev/docs/quickstart/)

## Run
```bash
minikube start
skaffold dev --port-forward
```
## Send query
```bash
curl -d '{"value":"hi"}' -H "Content-type: application/json" -X POST http://localhost:9000/v1/saga-workflow
```

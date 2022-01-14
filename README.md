# Saga POC

The POC demonstartes the approach to build micro service oriented arhitecture on the bases of the [gRPC](https://grpc.io/docs/languages/go/basics/) framework designed by Google. 

## Tools
### [gRPC](https://grpc.io/docs/languages/go/basics/)
### [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
### [buf](https://github.com/bufbuild/buf)
### [Skaffold](https://skaffold.dev/docs/quickstart/)

## Install etcd
```bash
docker run gcr.io/etcd-development/etcd@sha256:56aa454c329505b221216a60647ab315bf87c5bafe1ffe81c0af9014b496788a
```

## Run
```bash
minikube start
skaffold dev --port-forward
```
## Send query
```bash
curl -d '{"start":"start-value", "payload":"test-payload"}' -H "Content-type: application/json" -X POST http://localhost:9000/v1/state/init

curl -d '{"id": "368bc6e7-ad5c-4e39-b94a-11fe9bb3271e", "isRollback": false, "operation": {"from": "a", "to": "b", "name":"test"}, "payload":"v1"}' -H "Content-type: application/json" -X POST http://localhost:9000/v1/state/complete-operation
```

## Utilities 

### [gRPC curl](https://github.com/fullstorydev/grpcurl)


## Store 

### Send query
```bash
grpcurl -d '{"key":"test", "value":"test"}' -plaintext localhost:9005 storeapis.v1.StoreService/Put
```

### Service list
```bash
grpcurl -plaintext -import-path ./storeapis/v1 -proto storeapis.proto localhost:9005 list
```

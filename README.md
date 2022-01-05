# Saga POC

## Tools
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


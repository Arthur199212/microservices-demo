# Observability

## Using the Makefile

```bash
# Create the namespace
make create_namespace

# Deploy Jaeger operator
make add_jaeger_operator

# After the operator is deployed, create the Jaeger instance
make add_jaeger

# Finally, deploy the OpenTelemetry Collector
make add_otel_collector

# Clean resources
make clean
```

If you want to clean up after this, you can use the `make clean` to delete
all the resources created above. Note that this will not remove the namespace.
Because Kubernetes sometimes gets stuck when removing namespaces, please remove
this namespace manually after all the resources inside have been deleted,
for example with

```bash
kubectl delete namespaces observability
```

## Using minikube

```bash
# Create resources
minikube start
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.1/cert-manager.yaml
# Run from src/observability/ folder of this repository
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/jaeger-operator.yaml
kubectl apply -f k8s/jaeger.yaml
kubectl apply -f k8s/otel-collector.yaml
# Open port for otel-collector (http://localhost:30000)
kubectl port-forward -n observability svc/otel-collector 30000:4317
# Open port for Jaeger UI (http://localhost:16686)
kubectl port-forward -n observability svc/jaeger-query 16686:16686

# Delete resources one by one
kubectl delete -f k8s/otel-collector.yaml
kubectl delete -f k8s/jaeger.yaml
kubectl delete -f k8s/jaeger-operator.yaml
kubectl delete namespaces observability

# Delete the whole minikube cluster
minikube delete
```

JAEGER_OPERATOR_VERSION = v1.49.0

create_namespace:
	kubectl apply -f k8s/namespace.yaml

add_jaeger_operator:
	# kubectl apply -f k8s/jaeger-operator.yaml
	kubectl create -n observability -f https://github.com/jaegertracing/jaeger-operator/releases/download/$(JAEGER_OPERATOR_VERSION)/jaeger-operator.yaml

add_jaeger:
	kubectl apply -f k8s/jaeger.yaml

add_otel_collector:
	kubectl apply -f k8s/otel-collector.yaml

clean:
	- kubectl delete -f k8s/otel-collector.yaml
	- kubectl delete -f k8s/jaeger.yaml
	- kubectl delete -n observability -f https://github.com/jaegertracing/jaeger-operator/releases/download/$(JAEGER_OPERATOR_VERSION)/jaeger-operator.yaml
	- kubectl delete -f k8s/namespace.yaml

# Microservices Demo Application

| Service                           | Description                                                                                                                                 |
| --------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| [api-gateway](./src/api_gateway/) | Used as an entry point for clients to communicate with the rest of the application. It's built as RESP API and uses JSON for communication. |
| [cart](./src/products/)           | Stores the products and their quantity in the shopping cart by sessionId and allowes to retrieve them.                                      |
| [checkout](./src/products/)       | Retrieves user cart, prepares order and charges user provided card.                                                                         |
| [currency](./src/products/)       | Provides a list of supported currencies and makes a conversion from one currency to another.                                                |
| [payment](./src/products/)        | Charges provided debet/credit card with the given amount (mock) and return a transaction ID for reference.                                  |
| [products](./src/products/)       | Allows getting list of products and individual products by ID.                                                                              |
| [shipping](./src/products/)       | Provides shipping cost estimates based on the shopping cart (mock). Ships items to the given address and returns `tracingId` (mock).        |

## How To Run

### Prerequisites

- Docker for Desktop
- `kubectl`
- `skaffold`
- Minikube / kind

### Local Cluster

1. Launch a local k8s cluster `minikube start --cpus='2' --memory='3.9g'` or `kind create cluster`.
1. Verify the cluster is operating `minikube status` and it's possible to connect to control plane `kubectl get nodes`.
1. Run `skaffold run` or `skaffold dev`. This will build and deploy the application. If `skaffold dev` is used it will rebuild the images automatically as you change the code.
1. Run `kubectl get pods` to verify that the Pods are operating well.
1. Run `kubectl port-forward svc/api-gateway-service 8080:8080` to forward a port to the api-gateway service.
1. Use `localhost:8080` to access the api-gateway service.

### Cleanup

1. Use `skaffold delete` if you used `skaffold run` or just stop `skaffold dev` by pressing `ctrl+c`.
1. Run `minikube stop` to stop Minikube cluster or `minikube delete` to delete it.

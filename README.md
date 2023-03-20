# URL Shortener

A simple URL shortener service built with Go and Redis.

## Getting Started

1. To run this application, you will need to have Docker and Kubernetes installed on your machine
1. Clone the repository to your local machine using `git clone https://github.com/your-username/url-shortener-app.git`
1. Change directory to the project root using `cd url-shortener-app`
1. Build the Docker image using the command `docker build -t url-shortener:local-3 .`
1. Deploy the application to your Kubernetes cluster using the command `kubectl apply -f deployment.yml`
1. Access the application using the LoadBalancer service's external IP address. You can get the external IP address using the command `kubectl get svc url-shortener -n url-shortener-ns`
1. To use the NodePort service, get the NodePort using the command `kubectl get svc url-shortener-nodeport -n url-shortener-ns` and access the application using the Node's IP address and NodePort

## Local Development

To start the application locally, make sure you have Docker and Docker Compose installed, then run the following command:

```shell
docker-compose up
```


This will start both the API and Redis containers and bind the API to `http://localhost:8080`. You can test the service by sending a `POST` request with a JSON payload containing a `long_url` field to `http://localhost:8080`. The API will respond with a JSON payload containing the shortened URL in the `short_url` field.

## Deployment

To deploy the service to a Kubernetes cluster, make sure you have `kubectl` installed and configured to connect to the cluster, then run the following command:

```shell
kubectl apply -f deployment.yml
```


This will create a new namespace called `url-shortener-ns`, a Deployment with 3 replicas of the API container and a Redis container, a LoadBalancer Service for the API, and a NodePort Service for the API.

You can access the service by getting the external IP address of the LoadBalancer Service:

```shell
kubectl get svc -n url-shortener-ns
```


This will output the external IP address of the LoadBalancer Service. You can then send a `POST` request to `http://<external_ip>/` with a JSON payload containing a `long_url` field to create a new shortened URL, or send a `GET` request to `http://<external_ip>/<short_url>` to redirect to the original long URL.

## Environment Variables

The following environment variables can be used to configure the service:

- `REDIS_ADDR`: the address of the Redis server (default: `redis:6379`)

## Test
```shell
curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://www.google.com"}' http://localhost
ab -n 10000 -c 100 -p payload.json -T 'application/json' http://localhost/
```


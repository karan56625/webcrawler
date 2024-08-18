# Webcrawler and Webcrawler-Client

## Overview

Webcrawler: A Go application that crawls websites and generates a sitemap.
Webcrawler-Client: A Go client application that interacts with the Webcrawler service to initiate crawling and retrieve sitemaps.

## Prerequisites

### Go (Golang)
- Required for building and running the applications locally.
- Install from [Golang Official Site](https://go.dev/dl/).

### Docker
- Required for containerizing the applications.
- Install from [Docker Official Site](https://docs.docker.com/engine/install/).


### Kubernetes
- Required for deploying the applications on Kubernetes.
- Install kubectl from [Kubernetes Official Site](https://kubernetes.io/releases/download/).

### Access to a Kubernetes Cluster
- You need access to a Kubernetes cluster to deploy the applications. This can be a local cluster (like Minikube) or a managed Kubernetes service (like GKE, EKS, or AKS).

### Git
- Required for cloning the repositories.
- Install from Git Official Site.

## How to run

### Running Locally

#### Webcrawler

##### Clone the Repository

```bash
git clone https://github.com/yourusername/webcrawler.git
cd webcrawler
```
##### Install Dependencies

Ensure Go is installed, then run:

```bash
go mod download
```

##### Build the Project

Build the executable:

```bash
make build
```

##### Run the Application

Start the Webcrawler server:

```bash
./bin/webcrawler
```

Now, you can access the crawler at http://localhost:8081/crawl?url=<target-url>

There are few configuration that you can provide through environment variables.

**NUMBER_OF_WORKER** - To define how many workers should be used to crawl the url.

**WORKER_QUEUE_LENGTH** - To define the length of the worker queue. All crawled url are stored in a queue so that any worker can pick that from there.


#### Webcrawler-Client

##### Clone the Repository

```bash
git clone https://github.com/karan56625/webcrawler.git
cd webcrawler
```
##### Install Dependencies

Ensure Go is installed, then run:

```bash
go mod download
```

##### Build the Webcrawler Client

Build the executable:

```bash
make build-client
```

##### Run the Application
Make sure, Webcrawler server is running. You can run the webcrawler in separate terminal.

Run the Webcrawler client to crawl any URL:

```bash
./bin/webcrawler-client -url <target-url>
```
e.g.
```bash
./bin/webcrawler-client -url https://redhat.com
```

There are few configuration that you can provide through environment variables.

**WEBCRAWLER_HOST** - Host of the Webcrawler server. By default, it is http://localhost.

**WEBCRAWLER_PORT** - Port of the Webcrawler server. By default, it is 8081.


### Using Docker

#### Webcrawler

##### Build the Docker Image

Navigate to the Webcrawler project directory and build the webcrawler server image:

```bash
make docker-webcrawler
```
or 
```bash
docker build -f docker/webcrawler/Dockerfile -t webcrawler:1.0 .
```

##### Run the Docker Container

Run the container:

```bash
make docker-webcrawler-run
```
or
```bash
docker run -p 8081:8081 webcrawler:1.0
```

Now, you can access the crawler at http://localhost:8081/crawl?url=<target-url>

#### Webcrawler-Client

##### Build the Docker Image

Navigate to the Webcrawler project directory and build the webcrawler client image:

```bash
make docker-webcrawler-client
```
or
```bash
docker build -f docker/webcrawler-client/Dockerfile -t webcrawler-client:1.0 .
```

##### Run the container:

```bash
docker run --network="host" --rm webcrawler-client:1.0 -url <target-url>
```
e.g.
```bash
docker run --network="host" --rm webcrawler-client:1.0 -url https://redhat.com
```
Ensure to configure environment variables to connect to the Webcrawler service.


### Using Kubernetes

#### Webcrawler

Apply the Kubernetes configuration files:

```bash
kubectl apply -f k8s/webcrawler-deployment.yaml
kubectl apply -f k8s/webcrawler-service.yaml
```

##### Access the Service

If you are using an Ingress controller to manage external access, verify that it is correctly configured and functioning. Create the required Ingress to access the service. See: https://kubernetes.io/docs/concepts/services-networking/ingress/.

By default, `k8s/webcrawler-service.yaml` service of type LoadBalancer. You can change it based on the requirement. If you don't have configured load-balancer or Ingress controller, you might not able to access the service as there would be no external IP would be assigned.

In that case, you can access the service by doing the port-forwarding.

```bash
kubectl port-forward svc/webcrawler-service 8081:80
```
Now you can access the crawler at http://localhost:8081/crawl?url=<target-url>
Run the Webcrawler client to crawl any URL:
e.g.
```bash
./bin/webcrawler-client -url https://redhat.com
```

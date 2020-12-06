# urlshort
[![PkgGoDev](https://pkg.go.dev/badge/github.com/sachinkumarsingh092/urlshort)](https://pkg.go.dev/badge/github.com/sachinkumarsingh092/urlshort)
[![Go Report Card](https://goreportcard.com/badge/github.com/sachinkumarsingh092/urlshort)](https://goreportcard.com/report/github.com/sachinkumarsingh092/urlshort)
[![Build Status](https://travis-ci.com/sachinkumarsingh092/urlshort.svg?branch=master)](https://travis-ci.com/sachinkumarsingh092/urlshort)

Make your own short URLs. Use a customized YAML file to make the paths and URLs.

## Installation

  
- ### Cloning
  ```
  $ git clone https://github.com/sachinkumarsingh092/urlshort
  ```

- ### Verifing Checksums
  ```
  $ go mod verify
  ```

- ### Build locally
  
  - **Making binaries**
  ```
  $ go build -x main.go
  ```
  
  - **Running binaries**
  ```
  $ ./main [OPTIONS]
  ```
  
  The program will start on port 8080 locally.

  - **Test using `curl`**:
  ```
  curl -i localhost:8080/<SHORT_PATH>
  ```

- ### Build using Docker

  From the root of the directory:
  
  - **Building image from the Dockerfile**
  ```
  docker build -t <IMAGE_NAME>:latest .
  ```
  
  - **Running the container and port forwarding**
  ```
  docker run --publish 80:8080 -t <IMAGE_NAME>:latest 
  ```
  
  The host can now communicate with the container on port 80.

  - **Test using `curl`**:
  ```
  curl -i localhost:80/<SHORT_PATH>
  ```

  - **Push your image to dockerhub**
  ```
  docker push <YOUR_USERNAME>/urlshort:latest
  ```

## Deploying

We will use Kubernetes to manage the docker images which in turn will be deployed on the Google Kubernetes Engine (GKE). A GKE cluster is a managed set of Compute Engine virtual machines that operate as a single GKE cluster.

Go to the [official GKE documentation](https://cloud.google.com/kubernetes-engine) for more info about signing-in and billing details.

- ### Create a cluster
  We will use [Google Cloud SDK](https://cloud.google.com/sdk/docs/install) to create clusters:

  ```
  gcloud container clusters create urlshort \
    --num-nodes 1 \
    --enable-basic-auth \
    --issue-client-certificate \
    --zone <your-gcp-zone>
  ```

  Verify the access to cluster by using `kubectl`:

  ```
  kubectl get nodes
  ```

- ### Deploying to GKE
  To deploy your app to the GKE cluster you created, we need two Kubernetes objects.
  - A Deployment to define your app (/kubernetes/deployment.yaml).
  - A Service to define how to access your app (/kubernetes/service.yaml).

  ```
  kubectl apply -f deployment.yaml
  ```

  ```
  kubectl apply -f service.yaml
  ```

  Get the Service's external IP address:
  ```
  kubectl get services
  ```

  The `EXTERNAL-IP` can be then used to load the app in your web browsers.

## Options
```
-yaml=FILENAME
    Sets the YAML file to create URL map. (default "urls.yaml")
-help
    Print help.
```

## Documentations
Docs for handler package at [GoDoc](https://godoc.org/github.com/sachinkumarsingh092/urlshort/handler).

# Echo API

The Echo API is a versatile and very simple web server that provides an echo feature (echoes whatever you send to it). Under the hood, it implements a layered architecture on top of Echo Framework, GORM and SQLite.

## Table of Contents
- [Local prerequisites](#local-prerequisites)
- [Local execution](#local-execution)
- [Endpoints](#endpoints)

## <a name="local-prerequisites"></a>Local prerequisites

There are two ways to make this app run locally:
- Natively, using Golang 1.21.1 (Must be installed in your computer)
- Through Kubernetes, using `Docker Desktop`, its native Kubernetes cluster enabled and `kubectl`. (This way only tested in MacOS)

You'll find the necessary Kubernetes declarative files to make this app run in your local cluster under `resources/local-deployment/kubernetes/`

## <a name="local-execution"></a>Local execution

- Running this app natively:

    - From the root folder of this repo, execute `go mod tidy`.
    - Execute `go run cmd/main.go` to run the web server, which exposes the application through `localhost:8080`
    - Execute `go test ./...` to run unit tests.

- Running this app inside your local Kubernetes cluster:
    - From the root folder of this repo, execute `./deploy.sh`
    - Run option 1 to build the Docker image
    - Run option 2 to set the API configmap and launch the API pod (it exposes the app through `localhost:30001`)
    - Run option 3 whenever you've finished to terminate the pod and remove configmap.

## <a name="endpoints"></a>Endpoints

- `/v0/docs/` -> Swagger.
- `/v0/echo/` -> Creates an `echo` resource.
- `/v0/whats-echoed/` -> Returns `echo` instances created so far.
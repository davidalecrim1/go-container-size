# Container Size With Go
Comparing container size from "Hello World" code in Go on Dockerfile to a Multi Stage Dockerfile using Go.

## Getting Started

Building images:
```bash
docker build -f Dockerfile.big . -t web-app-1
docker build -f Dockerfile.medium . -t web-app-2
docker build -f Dockerfile.small . -t web-app-3
```

Running:
```bash
docker run -p 8080:8080 --name web-app-1 -d web-app-1
docker run -p 8080:8080 --name web-app-2 -d web-app-2
docker run -p 8080:8080 --name web-app-3 -d web-app-3
```

## Docker Images Results
```bash
REPOSITORY   TAG       IMAGE ID       CREATED      SIZE
web-app-1    latest    036861aef792   2 days ago   925MB
web-app-2    latest    913c022afc74   2 days ago   322MB
web-app-3    latest    f91bdec85028   2 days ago   4.92MB
```
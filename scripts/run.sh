#!/bin/sh

cd deployments/
docker compose exec -it -w /go/src/aquarium aquarium /bin/bash -c "go run /go/src/aquarium/main.go"
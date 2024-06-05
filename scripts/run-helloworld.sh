#!/bin/sh

docker exec -it -w /go/src/aquarium aquarium /bin/bash -c "go run /go/src/aquarium/main.go"
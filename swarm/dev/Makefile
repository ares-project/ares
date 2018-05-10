.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/ares-project/ares/cmd/swarm
	go build -o bin/geth     github.com/ares-project/ares/cmd/geth
	go build -o bin/bootnode github.com/ares-project/ares/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/ares-project/ares/swarm/...

GO_VERSION := 1.19  # <1>

KIND=${PWD}/kind
SERVICE=${PWD}/k8s

.PHONY: install-go init-go

setup: install-go init-go # <2>

install-go: # <3>
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go: # <4>
    echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
    echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

upgrade-go: # <5>
	sudo rm -rf /usr/bin/go
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

build:
	go build -o api cmd/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | \
	awk  '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html


#Setup cluster
.PHONY: create-cluster
create-cluster:
	echo 'Creating cluster ...'
	@cd ${KIND} && kind create cluster --config config.yaml
	@cd ${KIND} && kubectl apply -f ingress.yaml

.PHONY: local-service
local-service:
	@cd ${SERVICE} && kubectl apply -f deployment.yaml
	@cd ${SERVICE} && kubectl apply -f service.yaml

#Deletes cluster
.PHONY: clean
clean:
	@echo 'Deleting cluster...'
	kind delete cluster
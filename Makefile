GO_VERSION :=1.21

.PHONY: install-go init-go

setup: install-go init-go install-lint copy-hooks

install_go:
	 wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	 sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	 rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:                                                           
  echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
  echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

build:
	go build -o api cmd/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print ((int($$3) > 80) != 1)}'

report:
	go tool cover -html=coverage.out -o cover.html

check_format:
	test -z $$(go fmt ./...)

install-lint:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.56.2

 static-check:
	golangci-lint run

copy-hooks:
	chmod +x scripts/hooks/*
	cp -r scripts/hooks .git/.
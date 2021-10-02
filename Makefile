all: tidy vendor test build run

build:
	go build -o iptrans cmd/iptrans/main.go

run:
	./iptrans $(url)

clean:
	rm ./iptrans

tidy:
	go mod tidy

vendor:
	go mod vendor

.PHONY: coverage
coverage:
	go test \
		-race -covermode=atomic -timeout 30s \
		-coverprofile=coverage/coverage.out \
		github.com/captain-corgi/translate-host-to-ip/pkg/iptrans
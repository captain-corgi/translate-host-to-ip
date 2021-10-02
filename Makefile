all: tidy vendor coverage build run

build:
	go build -o ipconv cmd/iptrans/main.go

run:
	./ipconv $(url)

clean:
	rm ./ipconv
	rm -f vendor

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
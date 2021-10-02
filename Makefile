all: tidy vendor build run
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
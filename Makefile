all: tidy vendor build run

build:
	go build -o gos cmd/golang-os-example/main.go

run:
	./gos

clean:
	rm ./gos
	rm ./gos.exe
	rm -f vendor

tidy:
	go mod tidy

vendor:
	go mod vendor

.PHONY: coverage
# coverage:
# 	go test \
# 		-race -covermode=atomic -timeout 30s \
# 		-coverprofile=coverage/coverage.out \
# 		github.com/captain-corgi/golang-os-example/pkg/iptrans
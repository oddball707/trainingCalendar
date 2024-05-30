.PHONY: clean
clean:
	rm -f main main.zip

main:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ main.go

.PHONY: build
build: clean main

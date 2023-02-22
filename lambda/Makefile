.PHONY: clean
clean:
	rm -f main lambda.zip

main:
	go build -o $@ main.go

# compile the code to run in Lambda (local or real)
lambda:
	GOOS=linux GOARCH=amd64 go build -o lambda

.PHONY: build
build: clean lambda

lambda.zip: lambda
	zip $@ lambda
.PHONY: clean
clean:
	rm -f main lambda.zip

main:
	go build -o $@ main.go

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) main

.PHONY: build
build: clean lambda

lambda.zip: lambda
	zip $@ main
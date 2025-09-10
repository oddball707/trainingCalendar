build:
	docker build --file=./Dockerfile --rm=true -t training-cal .

run: build
	docker run -p 3000:8080 -d training-cal

run-dev:
	go run main.go

build-go:
	go mod download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /main .

build-npm:
	cd app;	npm install; npm run build;

npm:
	cd app; npm run dev;

cmd:
	go run main.go -cmd

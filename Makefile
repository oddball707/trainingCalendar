build:
	docker build --file=./Dockerfile --rm=true -t training-cal .

run:
	docker run -p 3000:8080 -d training-cal

build-go:
	go mod download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /main .

build-npm:
	cd app;	npm install; npm run build;

dev:
	go run main.go

npm:
	cd app; npm start;

heroku:
	git push heroku master

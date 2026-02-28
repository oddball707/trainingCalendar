build:
	docker build --file=./Dockerfile --rm=true -t training-cal .

run: build
	docker run -p 3000:8080 -d training-cal

run-dev:
	go run main.go

build-go:
	go mod download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./main .

build-npm:
	cd app;	npm install; npm run build;

npm:
	cd app; VITE_API_URL=http://localhost:8080 npm run dev;

cmd:
	go run main.go -cmd

test:
	gotestsum -- -v ./...

release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make release VERSION=v1.3.1"; \
		exit 1; \
	fi
	@echo "Creating and pushing tag $(VERSION)..."
	git tag $(VERSION)
	git push origin $(VERSION)
	@echo "Updating lambda/go.mod to use $(VERSION)..."
	@cd lambda && \
		sed -i.bak 's|github.com/oddball707/trainingCalendar v.*|github.com/oddball707/trainingCalendar $(VERSION)|' go.mod && \
		sed -i.bak '/^replace github.com\/oddball707\/trainingCalendar/d' go.mod && \
		rm go.mod.bak && \
		GOPROXY=direct go get github.com/oddball707/trainingCalendar@$(VERSION) && \
		go mod tidy
	@echo "Release $(VERSION) created successfully!"
	@echo "Next steps:"
	@echo "  1. Test the lambda build: cd lambda && go build"
	@echo "  2. Commit the updated lambda/go.mod: git add lambda/go.mod && git commit -m 'Update lambda to $(VERSION)'"
	@echo "  3. Push changes: git push"
	@echo "  4. Deploy via 'Update Lambda' GitHub Action"

dev-mode:
	@echo "Switching lambda to use local development mode..."
	@cd lambda && \
		sed -i.bak '/^replace github.com\/oddball707\/trainingCalendar/d' go.mod && \
		echo "" >> go.mod && \
		echo "replace github.com/oddball707/trainingCalendar => ../" >> go.mod && \
		rm go.mod.bak && \
		go mod tidy
	@echo "Lambda is now using local code via replace directive"

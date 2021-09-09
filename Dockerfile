FROM golang:1.16 as builder

RUN go get github.com/jordic/goics

WORKDIR /src/training
# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
COPY . .

# Fetch go modules
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/main ./main.go

FROM alpine:3.7
COPY --from=builder /bin/main /main
COPY --from=builder /src/training/data /data
EXPOSE 8080
ENTRYPOINT ["/main"]

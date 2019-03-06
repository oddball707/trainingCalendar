FROM golang:1.11.5 as builder

RUN go get github.com/jordic/goics

WORKDIR /src/training
# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# Fetch go modules
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

FROM alpine:3.7
COPY --from=builder /src/training /
EXPOSE 8080
ENTRYPOINT ["/main"]

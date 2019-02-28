FROM golang:1.11.5 as builder

RUN go get github.com/jordic/goics

WORKDIR /src/training
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

FROM alpine:3.7
COPY --from=builder /src/training /usr/local/bin
EXPOSE 8080
ENTRYPOINT main

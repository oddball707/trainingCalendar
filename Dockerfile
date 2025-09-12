FROM golang:1.25 as builder

ADD . /src
WORKDIR /src
# Force the go compiler to use modules
ENV GO111MODULE=on
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /main .

FROM node:24.7-alpine AS node_builder
COPY --from=builder /src/app ./
RUN npm install
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=builder /src/data /data
COPY --from=node_builder /dist ./web
RUN chmod +x ./main

EXPOSE 8080
CMD ./main

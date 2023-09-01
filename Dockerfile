#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/go-rest-template-app
COPY . .
RUN apk add --no-cache git
RUN go mod download
RUN go build -o app ./cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/go-rest-template-app/app /app
COPY --from=builder /go/src/go-rest-template-app/config/config.yml /config.yml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]
LABEL Name=go-rest-template-app
EXPOSE 8080
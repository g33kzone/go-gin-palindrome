 # Default to Go 1.11
 ARG GO_VERSION=1.15

 # First stage: build the executable.
 FROM golang:${GO_VERSION}-alpine AS builder

 RUN apk update && apk add git

 WORKDIR /app

 COPY go.mod .
 COPY go.sum .

 RUN go mod download

 COPY . .

 RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

 FROM scratch
 COPY --from=builder /app/go-gin-palindrome /app/
 EXPOSE 8080

 ENTRYPOINT [ "/app/go-gin-palindrome" ]
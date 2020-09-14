FROM golang:alpine as builder
RUN apk --no-cache add ca-certificates
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/squad-service /app/.env /app/
WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/squad-service"]
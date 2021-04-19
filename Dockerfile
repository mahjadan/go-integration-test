FROM golang:1.16.3-alpine3.13 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o go-app .

FROM alpine:3.11.3
COPY --from=builder /app/go-app .

EXPOSE 8080

ENTRYPOINT [ "./go-app" ]

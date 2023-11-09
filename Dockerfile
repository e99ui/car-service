FROM golang:alpine AS builder

RUN apk --no-cache add ca-certificates

LABEL stage=gobuilder

ENV CGO_ENABLED 0

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/main cmd/app/main.go


FROM scratch

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/main

CMD ["./main"]

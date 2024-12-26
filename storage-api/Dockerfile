FROM golang:1.23-alpine AS builder

WORKDIR /storage-api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o storageapi .

FROM alpine:latest

WORKDIR /storage-api

# AWS_ACCESS_KEY_ID=""
# AWS_SECRET_ACCESS_KEY=""
# AWS_DEFAULT_REGION=""
# RABBIT_MQ_URL =""

COPY --from=builder /storage-api/storageapi .
COPY --from=builder /storage-api/.env .

EXPOSE 8081

CMD [ "./storageapi" ]
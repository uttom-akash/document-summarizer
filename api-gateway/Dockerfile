FROM rust:1.72-slim AS builder

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y libssl-dev pkg-config && \
    cargo build --release && \
    apt-get clean && rm -rf /var/lib/apt/lists/*


FROM ubuntu:22.04

ENV STORAGE_API=http://storage-go-api:8081
ENV PROCESSOR_API=http://doc-processor-flask:5000

RUN apt-get update && \
    apt-get install -y libssl3 && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN adduser --disabled-password --gecos '' appuser

WORKDIR /app

COPY --from=builder /app/target/release/api-gateway .

RUN chown appuser:appuser api-gateway 

USER appuser

EXPOSE 3000

CMD [ "./api-gateway" ]
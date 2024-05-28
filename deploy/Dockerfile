FROM golang:1.20-alpine AS builder

WORKDIR /app
RUN apk add git && git clone https://github.com/icu-1/EasyImages.git .
RUN make build-linux

FROM alpine:3.19.0
WORKDIR /app
COPY --from=builder /app/bin/server ./server
COPY --from=builder /app/bin/config.yaml ./config.yaml
RUN chmod +x server

ENTRYPOINT ["sh","-c","./server"]
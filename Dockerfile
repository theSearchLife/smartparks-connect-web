FROM node:18.14.2-alpine3.17 AS npm-builder
WORKDIR /app
COPY ./public /app
WORKDIR /app
RUN npm i
RUN npm run build

FROM golang:1.19-buster AS go-builder
WORKDIR /go/src
COPY ./ .
RUN go build

FROM debian:buster-slim
# Required to send HTTPS requests to RockBLOCK
RUN apt-get update \
    && apt-get install --reinstall -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*
COPY --from=npm-builder /app/dist ./public/dist
COPY --from=go-builder /go/src/template_files ./template_files
COPY --from=go-builder /go/src/smartparks-connect-web .
ENTRYPOINT ["./smartparks-connect-web"]
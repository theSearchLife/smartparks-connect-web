FROM node:18.14.2-alpine3.17 AS npm-builder
WORKDIR /usr/app
COPY ./ /usr/app
WORKDIR /usr/app/public
RUN npm i
RUN npm run build

FROM golang:1.19-buster AS go-builder
WORKDIR /go/src
COPY ./ .
RUN go build

FROM debian:buster-slim
COPY --from=npm-builder /usr/app ./
COPY --from=go-builder /go/src/smartparks-connect-web .
ENTRYPOINT ["./smartparks-connect-web"]
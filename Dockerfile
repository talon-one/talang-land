FROM node:10.18.1-alpine AS client-builder

RUN apk add yarn

COPY ./client/package.json ./client/yarn.lock ./

RUN yarn

COPY ./client ./
RUN yarn generate

FROM golang:1.14-alpine AS server-builder

COPY ./server /go/src/github.com/talon-one/talang-land

ENV CGO_ENABLED=0

RUN go build -o /go/bin/talang-land github.com/talon-one/talang-land

FROM alpine:latest
COPY --from=server-builder /go/bin/talang-land /app/talang-land/talang-land
COPY --from=client-builder /dist /app/talang-land/static

WORKDIR /app/talang-land
CMD ./talang-land

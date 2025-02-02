FROM golang:1.23.5-alpine3.21 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./genproto ./genproto
COPY ./gateway ./gateway
COPY ./services ./services

COPY Makefile ./
RUN apk add --no-cache make
RUN make clean binaries

FROM alpine:3.21.2 AS gateway
WORKDIR /

COPY --from=build /app/bin/gateway ./
EXPOSE 8081
CMD [ "./gateway" ]

FROM alpine:3.21.2 AS authorisation
WORKDIR /

COPY --from=build /app/bin/authorisation ./
EXPOSE 9090
CMD [ "./authorisation" ]

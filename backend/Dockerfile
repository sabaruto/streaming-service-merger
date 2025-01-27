FROM golang:1.23.5-alpine3.21 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./genproto ./genproto
COPY ./services ./services

COPY Makefile ./
RUN apk add --no-cache make
RUN make clean binaries

FROM alpine:3.21.2 AS proxy
WORKDIR /

COPY --from=build /app/bin/proxy ./
EXPOSE 8081
CMD [ "./proxy" ]

FROM alpine:3.21.2 AS authorisation
WORKDIR /

COPY --from=build /app/bin/authorisation ./
EXPOSE 9090
CMD [ "./authorisation" ]

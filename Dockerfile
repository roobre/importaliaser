FROM golang:alpine as build

LABEL maintainer="Roberto Santalla <roobre@roobre.es>"


RUN mkdir -p /go/src/roob.re/importaliaser
WORKDIR /go/src/roob.re/importaliaser

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o importaliaser ./cmd


FROM alpine:latest

RUN mkdir -p /config
RUN mkdir -p /app
COPY --from=build /go/src/roob.re/importaliaser/importaliaser /app

ENTRYPOINT ["/app/importaliaser"]
CMD ["-j", "/config/store.json"]

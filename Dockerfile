FROM golang:1.18 AS build-env

WORKDIR /go/src/app

COPY go.mod .
COPY main.go .
COPY app ./app
COPY helpers ./helpers
COPY services ./services
COPY middlewares ./middlewares

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go mod tidy
RUN go vet -v
RUN go test -v

RUN go build -o /go/bin/app

# We don't use /base because we don't need OpenSSL, libSSL and glibc
FROM gcr.io/distroless/static

ARG CI_NAME=local
ENV CI_NAME=${CI_NAME}

ARG CI_VERSION=latest
ENV CI_VERSION=${CI_VERSION}

LABEL org.opencontainers.image.title="ReGraphQL"
LABEL org.opencontainers.image.description="A simple (yet effective) REST / HTTP to GraphQL router"
LABEL org.opencontainers.image.authors="ezequiel.aceto+regraphql@gmail.com"
LABEL org.opencontainers.image.url="https://hub.docker.com/repository/docker/eaceto/regraphql"
LABEL org.opencontainers.image.source="https://github.com/eaceto/ReGraphQL"
LABEL org.opencontainers.image.version="1.0.1"
LABEL dev.eaceto.regraphql.image.ci.name="${CI_NAME}"
LABEL dev.eaceto.regraphql.image.ci.version="${CI_VERSION}"
LABEL dev.eaceto.regraphql.health="http://localhost:8080/health/liveness"

COPY --from=build-env /go/bin/app /

EXPOSE 8080

HEALTHCHECK CMD curl --fail http://localhost:8080/health/liveness || exit 1

USER 1000
CMD ["/app"]
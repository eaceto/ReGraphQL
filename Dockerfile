FROM golang:1.18 AS build-env

WORKDIR /go/src/app

COPY app ./app
COPY helpers ./helpers
COPY go.mod .
COPY main.go .

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go mod tidy
RUN go vet -v
RUN go test -v

RUN go build -o /go/bin/app

# We don't use /base because we don't need OpenSSL, libSSL and glibc
FROM gcr.io/distroless/static

COPY --from=build-env /go/bin/app /

EXPOSE 8080

CMD ["/app"]
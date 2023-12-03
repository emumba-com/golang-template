FROM golang:1.19-alpine3.17

RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
      libressl \
      tar \
      git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

# Set the working directory
WORKDIR /go_service

# Copy the common source code to the container
COPY . ./

RUN go mod tidy
RUN GOOS=linux go build -a -tags musl .

# 2nd stage
FROM golang:1.17.7-alpine3.15
COPY --from=building-stage /go_service/gogintemplate /
EXPOSE 8000
CMD ["/gogintemplate"]

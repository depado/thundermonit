# Build Step
FROM golang:latest AS builder

# Prerequisites
RUN mkdir -p $GOPATH/src/github.com/Depado/thundermonit
ADD . $GOPATH/src/github.com/Depado/thundermonit
WORKDIR $GOPATH/src/github.com/Depado/thundermonit

# Build
ARG build
ARG version
RUN GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${version} -X main.Build=${build}" -o /tmp/thundermonit

# Final Step
FROM alpine

# Base packages
RUN apk update && apk upgrade && apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN rm -rf /var/cache/apk/*

# Copy binary from build step
COPY --from=builder /tmp/thundermonit /home/

# Define timezone
ENV TZ=Europe/Paris

# Define the ENTRYPOINT
WORKDIR /home
ENTRYPOINT ./thundermonit

# Document that the service listens on port 8080.
EXPOSE 8080
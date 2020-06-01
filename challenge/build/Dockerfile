# Builder
FROM golang:1.14 AS builder
LABEL "com.ekumenlabs.image.author.fullname"="Marcos Huck"
LABEL "com.ekumenlabs.image.author.email"="marcoshuck@ekumenlabs.com"

## Enable go modules
ENV GO111MODULE=on \
    CGO_ENABLED=1

## Download dependencies
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

## Build application
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o worker ./cmd/worker

## Move binary file to dist
WORKDIR /dist
RUN cp /build/worker .
RUN mkdir /data

#####################################################################

# Runner
FROM scratch
COPY --chown=0:0 --from=builder /dist /
COPY --chown=65534:0 --from=builder /data /data
USER 65534
WORKDIR /data
EXPOSE 3000
CMD ["/worker"]
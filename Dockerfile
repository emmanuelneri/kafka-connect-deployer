FROM golang:1.16-buster as builder
WORKDIR build
COPY . .
RUN rm -r examples
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags '-extldflags -s -w' -o /bin/app ./cmd

FROM scratch
COPY --from=builder /bin/app /bin/app
ENV KAFKA_CONNECT_URL "localhost:8083"
ENV CONNECTORS_FILES_DIR "/tmp"
ENV WAIT_START_TIME "1s"
ENV MAX_RETRY "3"
ENTRYPOINT ["/bin/app"]
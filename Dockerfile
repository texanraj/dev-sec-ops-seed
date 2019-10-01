FROM alpine:3

COPY bin/seed /seed

ENTRYPOINT ["/seed"]

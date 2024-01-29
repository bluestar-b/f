
FROM alpine:latest

RUN apk update && \
    apk add --no-cache \
    go \
    && rm -rf /var/cache/apk/*


WORKDIR /a
COPY *.go /a


EXPOSE 1234
RUN go mod init sus
RUN go build .

CMD ["./sus"]

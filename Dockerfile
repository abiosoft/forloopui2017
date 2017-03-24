FROM alpine:3.5

ADD app /app

EXPOSE 80 8000 8080

ENTRYPOINT ["/app"]

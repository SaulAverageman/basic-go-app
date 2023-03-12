FROM golang:1.19-alpine
RUN adduser -S goruntime
USER goruntime
WORKDIR /work
COPY ./src /work/
CMD ["go", "run", "/work/cmd/web/main.go"]
EXPOSE 8000
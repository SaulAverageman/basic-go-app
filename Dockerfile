FROM golang:1.19-alpine AS builder
RUN adduser -S goruntime
USER goruntime
WORKDIR /work
COPY ./src /work/
RUN cd cmd/web/ && CGO_ENABLED=0 go build -o /work/binary

FROM scratch
WORKDIR /work
COPY --from=builder /etc/passwd /etc/passwd
USER goruntime
COPY --from=builder /work /work
ENTRYPOINT [ "/work/binary" ]
EXPOSE 8000
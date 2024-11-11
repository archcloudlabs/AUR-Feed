FROM golang as builder
WORKDIR /go/
COPY . .

RUN CGO_ENABLED=0 go build

FROM alpine:latest
COPY --from=builder /go/AUR-Feed /opt/
ENTRYPOINT ["/opt/AUR-Feed"]

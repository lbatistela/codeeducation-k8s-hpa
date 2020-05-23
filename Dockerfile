FROM golang:1.14.2 as builder
COPY src /go/src
WORKDIR /go/src/squareroot
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o squareroot

FROM scratch
COPY --from=builder /go/src/squareroot/squareroot /
EXPOSE 8000
CMD ["/squareroot"]
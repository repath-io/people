FROM golang:1.15 as builder
ENV GO111MODULE on
COPY . .
RUN unset GOPATH
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service

FROM scratch
COPY --from=builder . .
ENTRYPOINT ["service"]
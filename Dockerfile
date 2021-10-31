FROM golang:1.17.2-alpine AS builder
WORKDIR /go/src
ADD . /go/src
#RUN go install golang.org/x/tools/cmd/goimports@latest
#RUN go mod vendor
RUN apk add make
RUN make build

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/entrypoint /app/
EXPOSE 80
ENTRYPOINT [ "/app/entrypoint"]
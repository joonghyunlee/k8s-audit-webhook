FROM golang:1.15

ENV GO111MODULE=on

WORKDIR ${GOPATH}/src/github.com/joonghyunlee/k8s-audit-webhook/
COPY . .

RUN go mod download
RUN go build

ENTRYPOINT ["./k8s-audit-webhook"]
EXPOSE 8080

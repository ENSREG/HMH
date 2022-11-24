FROM golang:1.19 AS build-env
RUN mkdir -p /src
WORKDIR /src
COPY main.go ./main.go
COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY test-server-0.1.0.tgz ./test-server-0.1.0.tgz
RUN go build main.go

RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 && \
chmod 700 get_helm.sh && \
./get_helm.sh

CMD ["./main"]
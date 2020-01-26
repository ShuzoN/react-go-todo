FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN apt-get update && apt-get install -y \
    vim \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/usr/local/go/bin/go", "run", "/go/src/app/src/main.go"]

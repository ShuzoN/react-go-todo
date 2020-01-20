FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN apt-get update && apt-get install -y \
    vim \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

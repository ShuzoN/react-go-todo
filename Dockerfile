FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN apt-get update && apt-get install -y \
    vim \
    lsof \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go get -d -v ./...
RUN go get -u github.com/oxequa/realize
RUN go install -v ./...

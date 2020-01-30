FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN apt-get update && apt-get install -y \
    vim \
    lsof \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN go get -d -v ./...
# https://github.com/oxequa/realize/issues/253
RUN go get gopkg.in/urfave/cli.v2@master
RUN go get github.com/oxequa/realize
RUN go get golang.org/x/tools/gopls
RUN go install -v ./...

FROM golang:1.9

ADD . /go/src/tests

WORKDIR /go/src/tests

RUN go get . && \
    go get github.com/mattes/migrate && \
    go get -u gopkg.in/alecthomas/gometalinter.v2 && \
    gometalinter.v2 --install && \
    curl -L "https://github.com/mattes/migrate/releases/download/v3.0.1/migrate.linux-amd64.tar.gz" | tar xvz

CMD ["sh", "-c", "tail -f /dev/null"]

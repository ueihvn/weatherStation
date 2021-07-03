FROM  golang:1.15.2-alpine3.12

WORKDIR $GOPATH/src/weatherStation/

ENV GOPATH /go

COPY . /go/src/weatherStation

RUN go install -v .

CMD [ "weatherStation" ]

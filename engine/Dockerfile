FROM golang

WORKDIR /go/src/engine
ADD . /go/src/engine
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/validator.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/unrolled/render
RUN go get github.com/thoas/stats
RUN go get github.com/codegangsta/negroni
RUN go get github.com/boltdb/bolt
RUN go get github.com/briandowns/stock-exchange/database
RUN go get github.com/pborman/uuid

RUN go install engine
EXPOSE 7777
ENTRYPOINT /go/bin/engine
FROM golang

WORKDIR /app

ADD . /app

RUN go get
RUN go install
RUN ls -lrt

EXPOSE 8080

CMD ["go", "run", "server.go"]
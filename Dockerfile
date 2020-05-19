FROM golang:1.13

WORKDIR /app

COPY . ./

RUN go get
RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

EXPOSE 8080

CMD ["/app/server"]
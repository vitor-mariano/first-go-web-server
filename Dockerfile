FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY . ./

RUN apk update
# RUN go install
RUN go build src/main.go

EXPOSE 3000
CMD ["./main"]

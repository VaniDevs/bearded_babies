FROM golang:alpine
RUN mkdir -p /app
WORKDIR /app
ADD . /app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]


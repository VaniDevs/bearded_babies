FROM golang:alpine
RUN mkdir -p /app
WORKDIR /app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get github.com/sendgrid/sendgrid-go
RUN go get github.com/appleboy/gin-jwt
RUN go get github.com/lib/pq

ADD . /app

RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]


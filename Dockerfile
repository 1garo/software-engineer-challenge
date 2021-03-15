FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/PicPay/software-engineer-challenge/
WORKDIR /go/src/github.com/PicPay/software-engineer-challenge
RUN go mod download
COPY . /go/src/github.com/PicPay/software-engineer-challenge
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/software-engineer-challenge github.com/PicPay/software-engineer-challenge

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
# RUN apk --update upgrade && \
#     apk add curl ca-certificates && \
#     update-ca-certificates 
COPY --from=builder /go/src/github.com/PicPay/software-engineer-challenge/build/software-engineer-challenge /usr/bin/software-engineer-challenge
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/software-engineer-challenge"]
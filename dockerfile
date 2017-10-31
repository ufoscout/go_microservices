#FROM golang:1.9 as builder
FROM instrumentisto/glide:0.13.0-go1.9 as builder

ENV GOPATH=/go

ENV SRC=$GOPATH/src/github.com/ufoscout/go_microservices
ADD ./ $SRC
WORKDIR $SRC

RUN glide up
RUN go build -o /hello ./web/hello.go

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.5  
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/

COPY --from=builder /hello /hello

CMD ["/hello"] 
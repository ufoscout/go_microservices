#FROM golang:1.9 as builder
FROM instrumentisto/glide:0.13.0-go1.9 as builder

ENV GOPATH=/go

ENV SRC=$GOPATH/src/github.com/ufoscout/go_microservices
ADD ./ $SRC
WORKDIR $SRC

RUN glide up
RUN go build -o /web ./web/main.go

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.5  
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.0.0/wait /wait
RUN chmod +x /wait

COPY --from=builder /web /web

CMD /wait && /web
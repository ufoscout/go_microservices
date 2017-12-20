#FROM golang:1.9-alpine as builder

FROM instrumentisto/dep:0.3-alpine as builder

ENV GOPATH=/go

ENV SRC=$GOPATH/src/github.com/ufoscout/go_microservices
ADD ./ $SRC
WORKDIR $SRC

RUN dep ensure
RUN go build -o /web ./web/main.go
#RUN go build -o /server ./server/main.go

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Build web image
FROM alpine:3.7

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.0.0/wait /wait
RUN chmod +x /wait

COPY --from=builder /web /web
RUN chmod +x /web

CMD /wait && /web
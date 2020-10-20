FROM golang:1.15-alpine as builder
 
WORKDIR /src/
COPY main.go .

RUN go build main.go
RUN chmod +x main

ENTRYPOINT [ "./main" ]

EXPOSE 80
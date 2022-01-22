FROM golang:latest
EXPOSE 8080

WORKDIR /usr/src/app
COPY ./ ./

RUN go build cmd/app/main.go

RUN apt-get update
RUN apt-get -y install postgresql-client

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

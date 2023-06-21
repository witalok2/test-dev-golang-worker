FROM golang:1.17 as build

ENV PORT 8081

WORKDIR /go/src/app

COPY . /go/src/app

RUN go build -o ./cmd/worker ./cmd/

ARG DATABASE_URL
ARG RABBITMQ_URI
ARG QUEUE_NAME

ENV GAE_ENV=production
ENV DATABASE_URL=$DATABASE_URL
ENV RABBITMQ_URI=$RABBITMQ_URI
ENV QUEUE_NAME=$QUEUE_NAME

CMD ["/go/src/app/cmd/worker"]

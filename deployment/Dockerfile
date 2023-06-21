FROM golang:1.17 as build

ARG GITHUB_TOKEN
ENV PORT 8081

WORKDIR /go/src/app

COPY . /go/src/app

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/YOUR-REPO".insteadOf "https://github.com/YOUR-REPO"
RUN export GOPRIVATE=github.com/YOUR-REPO/*

RUN go build -o ./cmd/worker ./cmd/

ARG DATABASE_URL
ARG RABBITMQ_URI
ARG QUEUE_NAME

ENV GAE_ENV=production
ENV DATABASE_URL=$DATABASE_URL
ENV RABBITMQ_URI=$RABBITMQ_URI
ENV QUEUE_NAME=$QUEUE_NAME

CMD ["/go/src/app/cmd/worker"]

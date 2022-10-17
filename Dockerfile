FROM golang:1.19.2-buster as go_app

FROM go_app as dev

COPY . /app

WORKDIR /app

# RUN go build -o /main

# CMD ["/main"]

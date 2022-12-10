FROM golang:1.19.2-buster as go_app

FROM go_app as base

COPY . /app/api-go

WORKDIR /app/api-go

# RUN go build -o /main

# CMD ["/main"]

# Setup buildpack
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://buildpack-registry.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

# Execute Buildpack
RUN STACK=heroku-20 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:20

COPY --from=build ./ /app/api-go
ENV HOME /app/api-go
WORKDIR /app/api-go
RUN useradd -m heroku
USER heroku
EXPOSE 80
CMD go run /app/api-go/main.go

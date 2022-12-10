# FROM golang:1.19.2-buster as go_app
FROM heroku/heroku:20-build as build

# FROM go_app as base

COPY . /app

WORKDIR /app

# RUN go build -o /main

# CMD ["/main"]

# Setup buildpack
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://buildpack-registry.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

# Execute Buildpack
RUN STACK=heroku-20 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:20

COPY --from=build ./ /app
ENV HOME /app
WORKDIR /app
RUN useradd -m heroku
USER heroku
EXPOSE 80
CMD /app/

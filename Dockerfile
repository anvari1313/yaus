FROM golang:1.13 AS build

ARG GIT_BRANCH
ARG GIT_SHA
ARG GIT_TAG
ARG BUILD_TIMESTAMP

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    TIMESTAMP=$BUILD_TIMESTAMP \
    CI_COMMIT_REF_SLUG=$GIT_BRANCH \
    CI_COMMIT_SHORT_SHA=$GIT_SHA

RUN mkdir -p /src

WORKDIR /src

COPY . /src

RUN alias go="http_proxy=$HTTP_PROXY go" && \
    make build-static-vendor && \
    mkdir -p /app && \
    cp ./yaus /app/

#
# 2. Runtime Container
#
FROM alpine:3.9

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

RUN apk add --update tzdata
RUN apk add --update ca-certificates
RUN apk add --update bash

RUN cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone && \
    mkdir -p /app && \
    mkdir -p /var/log && \
    chgrp -R 0 /var/log && \
    chmod -R g=u /var/log

WORKDIR /app

COPY --from=build /app /app

CMD ["./yaus", "serve"]

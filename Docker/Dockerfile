FROM alpine:3.7

RUN apk --no-cache --upgrade add ca-certificates \
    && update-ca-certificates --fresh \
    && apk --no-cache add --upgrade \
    git \
    make \
    go \
    curl \
    wget \
    maven \
    openjdk8 \
    less \
    py2-pip \
    && pip install --upgrade --no-cache-dir pip \
    && pip install --upgrade --no-cache-dir \
    awscli

FROM golang:1.11-alpine3.9

COPY py-requirements.txt .

RUN apk --no-cache --upgrade add ca-certificates \
    && update-ca-certificates --fresh \
    && apk --no-cache add --upgrade \
    coreutils \
    git \
    make \
    docker \
    curl \
    wget \
    maven \
    openjdk8 \
    less \
    tree \
    groff \
    python3 \
    binutils \
    gcc \
    musl-dev \
    jq \
    && apk --no-cache add --upgrade --virtual .build-dependencies \
    python3-dev \
    libffi-dev \
    openssl-dev \
    && pip3 install --upgrade --no-cache-dir pip \
    && pip3 install --upgrade --no-cache-dir -r py-requirements.txt \
    && rm -rf /tmp/** \
    && apk del .build-dependencies

COPY dockerd-start.sh /usr/local/bin/

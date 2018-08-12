FROM alpine:3.7

ENV GOPATH="/go" \
    DEP_VERSION="0.4.1" \
    DEP_BINARY="dep-linux-amd64"

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
    musl-dev \
    && pip install --upgrade --no-cache-dir pip \
    && pip install --upgrade --no-cache-dir \
    awscli \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && wget "https://github.com/golang/dep/releases/download/v$DEP_VERSION/$DEP_BINARY" -q -O "$GOPATH/bin/dep" \
    && chmod +x "$GOPATH/bin/dep"

ENV PATH="$GOPATH/bin:/usr/local/go/bin:$PATH"

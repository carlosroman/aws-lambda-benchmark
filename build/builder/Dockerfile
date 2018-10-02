FROM alpine:3.8

ENV GOPATH="/go" \
    DEP_VERSION="0.4.1" \
    DEP_BINARY="dep-linux-amd64"

COPY py-requirements.txt .

RUN apk --no-cache --upgrade add ca-certificates \
    && update-ca-certificates --fresh \
    && apk --no-cache add --upgrade \
    coreutils \
    git \
    make \
    docker \
    go \
    curl \
    wget \
    maven \
    openjdk8 \
    less \
    tree \
    groff \
    py2-pip \
    python2-dev \
    musl-dev \
    && pip install --upgrade --no-cache-dir pip \
    && pip install --upgrade --no-cache-dir -r py-requirements.txt \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && wget "https://github.com/golang/dep/releases/download/v$DEP_VERSION/$DEP_BINARY" -q -O "$GOPATH/bin/dep" \
    && chmod +x "$GOPATH/bin/dep"

COPY dockerd-start.sh /usr/local/bin/

ENV PATH="$GOPATH/bin:/usr/local/go/bin:$PATH"

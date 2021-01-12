ARG GO_VERSION=1.15.3

FROM golang:${GO_VERSION}-alpine

ENV PROTOBUF_VERSION="3.13.0"
ENV PROTOBUF_ZIP=protoc-${PROTOBUF_VERSION}-linux-x86_64.zip
ENV PROTOBUF_URL=https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/${PROTOBUF_ZIP}

ENV GLIBC_VERSION="2.23-r3"
ENV ALPINE_GLIBC_URL=https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/
ENV GLIBC_PKG=glibc-${GLIBC_VERSION}.apk
ENV GLIBC_BIN_PKG=glibc-bin-${GLIBC_VERSION}.apk


RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN apk add --update --no-cache -t deps wget ca-certificates \
  && cd /tmp \
  # install glibc
  && wget ${ALPINE_GLIBC_URL}${GLIBC_PKG} \
  && wget ${ALPINE_GLIBC_URL}${GLIBC_BIN_PKG} \
  && apk add --allow-untrusted ${GLIBC_PKG} ${GLIBC_BIN_PKG} \
  && /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc-compat/lib/ \
  && echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf \
  # install protobuf
  && wget ${PROTOBUF_URL} \
  && unzip ${PROTOBUF_ZIP} 'bin/*' -d /usr \
  # Cleanup
  && apk del --purge deps \
  && rm -rf /tmp/* /var/cache/apk/*

# libstdc++ is required for running protoc binary
RUN apk add --update --no-cache libstdc++ \
  && rm -rf /var/cache/apk/*

RUN apk --no-cache add --virtual .scm git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build
RUN go get -u github.com/golang/protobuf/protoc-gen-go

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go generate

COPY . .
RUN go build -o main .

WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .


EXPOSE 8080

CMD ["/dist/main", "--server","blackstone:9999"]

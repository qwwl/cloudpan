FROM golang:alpine
MAINTAINER alphayan "alphayyq@163.com"
WORKDIR /build
COPY . /build
ENV CGO_ENABLED=0
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add upx \
    && rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $HOME/.cache \
    && go build -mod=vendor -ldflags '-w -s -extldflags "-static"' -o cloudpan \
    && upx -9 cloudpan \
	&& cp cloudpan /run \
	&& cp config.toml /run
FROM alpine
MAINTAINER alphayan "alphayyq@163.com"
COPY --from=0 /run /run
WORKDIR /run
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache ca-certificates \
    && apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $HOME/.cache
CMD ["./cloudpan"]

FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY bili-danmu-auth/api/etc /usr/local/bin/etc
RUN go build -ldflags='-s -w -extldflags "-static"' -tags osusergo,netgo,sqlite_omit_load_extensio -o /usr/local/bin/danmu-auth-api bili-danmu-auth/api/danmu-auth.go


# FROM scratch
FROM alpine as final

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /usr/local/bin/danmu-auth-api /usr/local/bin/danmu-auth-api
COPY --from=builder /usr/local/bin/etc /usr/local/bin/etc

EXPOSE 8888

CMD [ "/usr/local/bin/danmu-auth-api", "-f", "/usr/local/bin/etc/danmu-auth.yaml" ]
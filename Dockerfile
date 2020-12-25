# docker 编译时 go get 失败解决方法
# https://blog.csdn.net/asty9000/article/details/107720900
# https://www.cnblogs.com/zhangmingcheng/p/12294156.html

# 1. 准备基础镜像
FROM golang:1.15.6-alpine as skinshub-api-base
WORKDIR /opt/app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct

# 2. 下载依赖
FROM skinshub-api-base as skinshub-api-deps
COPY go.mod go.sum ./
# download deps
RUN go mod download

# 3. 编译
FROM skinshub-api-deps as skinshub-api-build
COPY . .
# build
RUN go build -o main

# 4. 生产镜像
FROM scratch
WORKDIR /opt/app
COPY --from=skinshub-api-build /opt/app .
EXPOSE 1323
CMD ["./main"]
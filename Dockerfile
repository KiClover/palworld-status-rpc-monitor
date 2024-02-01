FROM golang:1.21.6-alpine3.17 as builder

WORKDIR /build
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -ldflags="-s -w" -o /build/${PROJECT}_rpc ${PROJECT}.go \

# Define the project name | 定义项目名称
ARG PROJECT=palworldmonitor
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=palworldmonitor.yaml
# Define the author | 定义作者
ARG AUTHOR="kiclover@email.cn"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY ./${PROJECT}_rpc ./
COPY ./etc/${CONFIG_FILE} ./etc/

EXPOSE 8080

ENTRYPOINT ./${PROJECT}_rpc -f etc/${CONFIG_FILE}

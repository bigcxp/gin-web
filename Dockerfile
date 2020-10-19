FROM golang:1.14 AS gin-web

RUN echo "----------------- 后端Gin Web构建(Production) -----------------"
# 环境变量
# 开启go modules
ENV GO111MODULE=on
# 使用国内代理, 避免被墙资源无法访问
ENV GOPROXY=https://goproxy.cn
# 定义应用运行目录
ENV APP_HOME /app/gin-web-prod

RUN mkdir -p $APP_HOME

# 设置运行目录
WORKDIR $APP_HOME

# 这里的根目录以docker-compose.yml配置build.context的为准
# 拷贝宿主机go.mod / go.sum文件到当前目录
COPY ./gin-web/go.mod ./gin-web/go.sum ./
# 下载依赖文件
RUN go mod download

# 拷贝宿主机全部文件到当前目录
COPY ./gin-web .

# 通过packr2将配置文件写入二进制文件
# 构建packr2
RUN cd $GOPATH/pkg/mod/github.com/gobuffalo/packr/v2@v2.8.0/packr2 && go build && chmod +x packr2
# 回到app目录运行packr2命令
RUN cd $APP_HOME && $GOPATH/pkg/mod/github.com/gobuffalo/packr/v2@v2.8.0/packr2/packr2 build

# 构建应用
RUN go build -o main-prod .

# mysqldump需要一些依赖库这里直接使用alpine-glibc
# alpine镜像瘦身
FROM frolvlad/alpine-glibc:alpine-3.12

# 定义程序运行模式
ENV GIN_WEB_MODE production
# 定义应用运行目录
ENV APP_HOME /app/gin-web-prod

RUN mkdir -p $APP_HOME

# 设置运行目录
WORKDIR $APP_HOME

COPY --from=gin-web $APP_HOME/main-prod .
COPY --from=gin-web $APP_HOME/gitversion .

# 拷贝mysqldump文件(binlog刷到redis会用到)
COPY ./gin-web/docker-conf/mysql/mysqldump /usr/bin/mysqldump

# alpine时区修改
# apk仓库使用国内源
# 设置时区为上海
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update \
  && apk add tzdata \
  && apk add libstdc++ \
  && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo "Asia/Shanghai" > /etc/timezone
# 验证时区是否已修改
# RUN date -R

# 暴露端口
EXPOSE 8080

# 启动应用(daemon off后台运行)
CMD ["./main-prod", "-g", "daemon off;"]

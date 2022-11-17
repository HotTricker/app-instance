# app-instance

Minimal back-end app using golang

使用gin-gonic框架，使用yaml配置文件，logrus+(file-rotatelogs)管理日志

## 私有仓库使用：

1. 环境变量 GOPRIVATE 用来控制 go 命令把哪些仓库看做是私有的仓库，跳过 proxy server 和校验检查

    `go env -w GOPRIVATE=192.168.33.236`

2. gitlab如果不是https的，要加上环境变量GOINSECURE

    `go env -w GOINSECURE=192.168.33.236`

## redis：

一主一从三哨兵

若主未能发现从，有可能是redis无法用ip访问（127.0.0.1可以访问），注释掉NETWORK下面的bind并设置protected-mode 为 no
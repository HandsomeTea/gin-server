# gin-server
common http server by go

## 入门步骤
- 项目创建和运行
  - 创建工作区：go work init /xxx/go
  - cd /xxx/go
  - mkdir src
  - clone当前项目到src目录下
  - cd gin-server
  - go mod init gin-server
  - go get -u github.com/gin-gonic/gin
  - go run main.go

- 已有项目安装依赖和运行
  - go mod tidy
  - go run main.go

## 开发相关
- 使用vscode，安装插件`Go`。
- 热加载：
  - 安装`air`：`go install github.com/air-verse/air@latest`。
  - 检查`air`是否安装成功：`air -v`，不成功一般是因为`air`安装在了go的工作区，把工作区中的`air`做软连接。
  - 运行：`air`。

## 网关配置
- 接口网关：需要配置环境变量`DEPLOY_MODE`为`api-gateway`。
- 使用网关辅助实现A/B测试：需要配置环境变量`DEPLOY_MODE`为`ab-test`。
  - 流量分配百分比：配置环境变量`AB_TEST_DIVERSION_PERCENTAGE`，默认值为`0`，取值范围`0-100`。
  - 流量分配：配置环境变量`AB_TEST_MODE`：
    - `redirect`：对`/`请求进行一次处理，根据`AB_TEST_DIVERSION_PERCENTAGE`配置返回`301`状态的http请求，接受端收到该请求自行处理重定向；返回`200`状态不做处理。
    - `proxy`：对`/index.html`请求进行一次处理，根据`AB_TEST_DIVERSION_PERCENTAGE`配置返回`305`状态的http请求，接受端收到该请求自行处理代理；返回`200`状态不做处理。
- `DEPLOY_MODE`其它值配置则接口网关和A/B测试同时支持。

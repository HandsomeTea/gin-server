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

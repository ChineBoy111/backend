# gin 集成 swagger

### 安装 swagger

```shell
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

export GOBIN="$HOME/bin"
go install github.com/swaggo/swag/cmd/swag@latest
```

### 初始化 swagger

```shell
cd src
swag init # 将在 src 目录下创建 docs 目录
```

### swagger 注释

项目注释 [main.go](../main.go)

```go
// @title       gin-gorm
// @version     0.0.1
// @description {"go": ["gin", "gorm"], "typescript": ["vue3", "vite"]}
func main() {
	defer cmd.Done()
	cmd.Start()
}
```

api 注释

- path: URL 参数
- formData: 请求体参数

```go
// @Tag         标签
// @Summary     api 功能，简略
// @Description api 功能，详细
// @Accept      json
// @Produce     json
// @Param       参数名   path/formData   数据类型   是否必须   描述
// @Success     成功时的响应状态码   {string/object}   数据类型   描述
// @Failure     失败时的响应状态码   {string/object}   数据类型   描述
// @Router      路由 [请求方法]

// 例
// @Tag         用户 api
// @Summary     用户登录，简略
// @Description 用户登录，详细
// @Accept      json
// @Produce     json
// @Param       username   formData   string   true   "用户名"
// @Param       password   formData   string   true   "密码"
// @Success     200   {string}   string   "登录成功"
// @Failure     401   {string}   string   "登录失败"
// @Router      /api/v1/public/user/login [post]
```


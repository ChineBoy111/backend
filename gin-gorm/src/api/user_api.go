package api

import (
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/service/dto"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// Login api 的 swagger 注释
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
func (UserApi) Login(ctx *gin.Context) {
	//? ctx.AbortWithStatusJSON(http.StatusOK /* 200 */, gin.H{
	//? 	"msg": "Login ok",
	//? } /* gin.H 是 map[string]any 的别名 */)

	var userLoginDto dto.UserLoginDto
	//! ctx.ShouldBind 检查请求方式 GET, POST, ... 和 Content-Type 以自动解析并绑定
	//! 例如 "application/json" -> JSON 绑定，"application/xml" -> XML 绑定
	//! ctx.ShouldBind 与 ctx.Bind 相似，不同的是
	//! 绑定失败时，ctx.ShouldBind 不会将响应状态码设置为 404 或终止
	err := ctx.ShouldBind(&userLoginDto) // 自动解析并绑定
	if err != nil {
		global.Logger.Errorln(err.Error())
		ClientErr(ctx, Resp{
			Msg: err.Error(),
		})
	}

	Ok(ctx, Resp{
		Data: userLoginDto,
	})
}

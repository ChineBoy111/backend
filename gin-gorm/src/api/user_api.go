package api

import (
	"net/http"

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
	//// ctx.AbortWithStatusJSON(http.StatusOK /* 200 */, gin.H{
	//// 	"msg": "Login ok",
	//// } /* gin.H 是 map[string]any 的别名 */)

	ctx.AbortWithStatusJSON(http.StatusOK, Resp{
		Msg: "Login ok",
	})
}

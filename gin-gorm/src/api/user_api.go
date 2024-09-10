package api

import (
	"bronya.com/gin-gorm/src/dto"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/service"
	"bronya.com/gin-gorm/src/util"
	"github.com/gin-gonic/gin"
)

const (
	USER_LOGIN_ERR = iota + 1000
	USER_INSERT_ERR
	USER_SELECT_BY_ID_ERR
	USER_SELECT_BY_PAGE_ERR
	USER_UPDATE_ERR
)

type UserApi struct {
	UserService *service.UserService //! 组合 UserService
}

// ! UserApi 单例
var userApi *UserApi

func NewUserApi() *UserApi {
	if userApi == nil {
		userApi = &UserApi{
			UserService: service.NewUserService(),
		}
	}
	return userApi
}

// UserLogin api 的 swagger 注释
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
func (userApi UserApi) UserLogin(ctx *gin.Context) { //! 不使用指针接收
	//// ctx.AbortWithStatusJSON(http.StatusOK /* 200 */, gin.H{
	//// 	   "msg": "SelectUserByUsernameAndPassword ok",
	//// } /* gin.H 是 map[string]any 的别名 */)

	var userLoginDto dto.UserLoginDto
	//* ctx.ShouldBind 检查请求方式 GET, POST, ... 和 Content-Type 以自动解析并绑定
	//* 例如 "application/json" -> json 绑定，"application/xml" -> xml 绑定
	//* ctx.ShouldBind 与 ctx.Bind 相似，不同的是
	//* 绑定失败时，ctx.ShouldBind 不会将响应状态码设置为 404 或终止
	validationErrs := ctx.ShouldBind(&userLoginDto) //! 自动解析并绑定
	if validationErrs != nil {
		global.Logger.Errorln(validationErrs.Error())
		Err(ctx, Resp{
			Code: USER_LOGIN_ERR,
			Msg:  validationErrs.Error(),
		})
		return
	}

	user, err := userApi.UserService.SelectUserByUsernameAndPassword(&userLoginDto)
	if err != nil {
		Err(ctx, Resp{
			Code: USER_LOGIN_ERR,
			Msg:  err.Error(),
		})
		return
	}

	token, _ := util.GenToken(user.ID, user.Username)
	Ok(ctx, Resp{
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func (userApi UserApi) InsertUser(ctx *gin.Context) {
	var userInsertDto dto.UserInsertDto
	//* ctx.ShouldBind(any)
	validationErrs := ctx.ShouldBind(&userInsertDto)
	if validationErrs != nil {
		global.Logger.Errorln(validationErrs.Error())
		Err(ctx, Resp{
			Code: USER_INSERT_ERR,
			Msg:  validationErrs.Error(),
		})
		return
	}

	err := userApi.UserService.InsertUser(&userInsertDto)
	if err != nil {
		Err(ctx, Resp{
			Code: USER_INSERT_ERR,
			Msg:  err.Error(),
		})
		return
	}

	Ok(ctx, Resp{
		Data: userInsertDto,
	})
}

func (userApi UserApi) SelectUserById(ctx *gin.Context) {
	var commonIdDto dto.IdDto
	//* ctx.ShouldBindUri(any)
	validationErrs := ctx.ShouldBindUri(&commonIdDto)
	if validationErrs != nil {
		global.Logger.Errorln(validationErrs.Error())
		Err(ctx, Resp{
			Code: USER_SELECT_BY_ID_ERR,
			Msg:  validationErrs.Error(),
		})
		return
	}
	user, err := userApi.UserService.SelectUserById(&commonIdDto)
	if err != nil {
		Err(ctx, Resp{
			Code: USER_SELECT_BY_ID_ERR,
			Msg:  err.Error(),
		})
		return
	}
	Ok(ctx, Resp{
		Data: user,
	})
}

func (userApi UserApi) SelectPaginatedUser(ctx *gin.Context) {
	var paginateDto dto.PaginateDto
	//* ctx.ShouldBind(any)
	validationErrs := ctx.ShouldBind(&paginateDto)
	if validationErrs != nil {
		global.Logger.Errorln(validationErrs.Error())
		Err(ctx, Resp{
			Code: USER_SELECT_BY_PAGE_ERR,
			Msg:  validationErrs.Error(),
		})
		return
	}
	userArr, total, err := userApi.UserService.SelectPaginatedUser(&paginateDto)
	if err != nil {
		Err(ctx, Resp{
			Code: USER_SELECT_BY_PAGE_ERR,
			Msg:  err.Error(),
		})
		return
	}
	Ok(ctx, Resp{
		Data:  userArr,
		Total: total,
	})
}

func (userApi UserApi) UpdateUser(ctx *gin.Context) {
	var userUpdateDto dto.UserUpdateDto
	validationErrs := ctx.ShouldBindUri(&userUpdateDto)
	if validationErrs != nil {
		global.Logger.Errorln(validationErrs.Error())
		Err(ctx, Resp{Code: USER_UPDATE_ERR, Msg: validationErrs.Error()})
		return
	}
	validationErrs = ctx.ShouldBind(&userUpdateDto)
	if validationErrs != nil {
		global.Logger.Errorln(validationErrs.Error())
		Err(ctx, Resp{Code: USER_UPDATE_ERR, Msg: validationErrs.Error()})
		return
	}
	err := userApi.UserService.UpdateUser(&userUpdateDto)
	if err != nil {
		Err(ctx, Resp{
			Code: USER_UPDATE_ERR,
			Msg:  err.Error(),
		})
		return
	}
	Ok(ctx, Resp{
		Msg: "Update ok",
	})
}

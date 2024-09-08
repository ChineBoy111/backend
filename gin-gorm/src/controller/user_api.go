package controller

import (
	"bronya.com/gin-gorm/src/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

type UserApi struct {
}

var userApi *UserApi

func NewUserApi() *UserApi {
	if userApi == nil {
		userApi = &UserApi{}
	}
	return userApi
}

// Login controller 的 swagger 注释
// @Tag         用户 controller
// @Summary     用户登录，简略
// @Description 用户登录，详细
// @Accept      json
// @Produce     json
// @Param       username   formData   string   true   "用户名"
// @Param       password   formData   string   true   "密码"
// @Success     200   {string}   string   "登录成功"
// @Failure     401   {string}   string   "登录失败"
// @Router      /controller/v1/public/user/login [post]
func (userApi UserApi) Login(ctx *gin.Context) {
	// ctx.AbortWithStatusJSON(http.StatusOK /* 200 */, gin.H{
	// 	   "msg": "Login ok",
	// } /* gin.H 是 map[string]any 的别名 */)

	var userLoginDto UserLoginDto
	//! ctx.ShouldBind 检查请求方式 GET, POST, ... 和 Content-Type 以自动解析并绑定
	//! 例如 "application/json" -> JSON 绑定，"application/xml" -> XML 绑定
	//! ctx.ShouldBind 与 ctx.Bind 相似，不同的是
	//! 绑定失败时，ctx.ShouldBind 不会将响应状态码设置为 404 或终止
	validationErrs := ctx.ShouldBind(&userLoginDto) // 自动解析并绑定
	if validationErrs != nil {
		global.Logger.Errorln(validationErrs.Error())
		ClientErr(ctx, Resp{
			//! 响应可读的错误消息
			Msg: parseValidationErrors(validationErrs.(validator.ValidationErrors), &userLoginDto).Error(),
		})
	}

	Ok(ctx, Resp{
		Data: userLoginDto,
	})
}

// parseValidationErrors 响应可读的错误消息
func parseValidationErrors(validationErrs validator.ValidationErrors, dtoPtr any) error {
	//! 通过反射，获取指针 dtoPtr 指向的元素的类型
	dtoType := reflect.TypeOf(dtoPtr).Elem()
	log.Println("dtoTypeName =", dtoType.Name()) //! dtoTypeName = dto.UserLoginDto

	var errMsgArr []string
	for _ /* idx */, validationErr := range validationErrs /* validator.ValidationErrors */ {
		//! 获取校验出错的字段名 fieldName
		fieldName := validationErr.Field()
		// 校验出错的字段名 Username
		log.Println("fieldName =", fieldName) //* fieldName = Username

		//! 获取校验出错的错误标签名 errTagName
		errTagName := validationErr.Tag()
		// 校验出错的错误标签名 not_admin
		log.Println("errTagName =", errTagName) //* errTagName = not_admin

		//! 获取指定字段名 fieldName 对应的结构化字段 structField
		structField, _ := dtoType.FieldByName(fieldName)

		//! 指定错误消息标签名 errMsgTagName 为 `${errTagName}_err`
		errMsgTagName := fmt.Sprintf("%s_err", errTagName)
		log.Println("errMsgTagName =", errMsgTagName) //* errMsgTagName = not_admin_err

		//! 获取错误消息标签的值，即可读的错误消息 errMsg
		errMsg /* errMsgTagValue */ := structField.Tag.Get(errMsgTagName)
		if errMsg == "" {
			errMsg = structField.Tag.Get("msg")
		}
		if errMsg == "" {
			errMsg = fmt.Sprintf("%s: %s error", fieldName, errTagName) // Username: not_admin error
		}
		log.Println("err_msg =", errMsg) //! errMsg = 用户名不能为空
		errMsgArr = append(errMsgArr, errMsg)
	}
	return errors.New(strings.Join(errMsgArr, "; "))
}

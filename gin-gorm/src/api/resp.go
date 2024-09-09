package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Resp struct {
	//! -           json 编解码时，忽略该字段
	Status int `json:"-"`
	//! code        json 中的字段名为 code
	//! omitempty   如果该字段为空，则 json 编解码时，忽略该字段
	Code int `json:"code,omitempty"`
	//  msg         json 中的字段名为 msg
	Msg string `json:"msg,omitempty"`
	//  model        json 中的字段名为 model
	Data any `json:"model,omitempty"`
}

func (resp Resp) IsEmpty() bool {
	return reflect.DeepEqual(resp, Resp{}) //? 通过反射，判断 resp 是否为空
}

func Ok(ctx *gin.Context, resp Resp) { //* 2xx 成功
	if resp.IsEmpty() {
		ctx.AbortWithStatus(http.StatusOK) // 200
		return
	}
	if resp.Status < 200 || resp.Status >= 300 {
		resp.Status = http.StatusOK // 200
	}
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func ClientErr(ctx *gin.Context, resp Resp) { //* 4xx 客户端错误
	if resp.IsEmpty() {
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}
	if resp.Status < 400 || resp.Status >= 500 {
		resp.Status = http.StatusBadRequest // 400
	}
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func ServerErr(ctx *gin.Context, resp Resp) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
	}
	if resp.Status < 500 {
		resp.Status = http.StatusInternalServerError // 500
	}
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

// ParseValidationErrors 响应可读的错误消息
func ParseValidationErrors(validationErrs validator.ValidationErrors, dtoPtr any) error {
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

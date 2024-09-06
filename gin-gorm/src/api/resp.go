package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	//! -           json 编解码时，忽略该字段
	Status int `json:"-"`
	//! code        json 中的字段名为 code
	//! omitempty   如果该字段为空，则 json 编解码时，忽略该字段
	Code int `json:"code,omitempty"`
	//* msg         json 中的字段名为 msg
	Msg string `json:"msg,omitempty"`
	//* data        json 中的字段名为 data
	Data any `json:"data,omitempty"`
}

func (resp Resp) IsEmpty() bool {
	return reflect.DeepEqual(resp, Resp{}) //? 反射，判断 resp 是否为空
}

func Ok(ctx *gin.Context, resp Resp) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(http.StatusOK)
		return
	}

	if resp.Status < 200 || resp.Status >= 300 {
		resp.Status = http.StatusOK // 200
	}
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func Err(ctx *gin.Context, resp Resp) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}
	if resp.Status < 400 || resp.Status >= 500 {
		resp.Status = http.StatusBadRequest // 400
	}
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

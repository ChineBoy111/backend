package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Resp struct {
	//! -           json 编解码时，忽略该字段
	Status int `json:"-"`
	//! code        json 中的字段名为 code
	//! omitempty   如果该字段为空，则 json 编解码时，忽略该字段
	Code int `json:"code,omitempty"`
	//  msg         json 中的字段名为 msg
	Msg string `json:"msg,omitempty"`
	//  data        json 中的字段名为 data
	Data any `json:"data,omitempty"`
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

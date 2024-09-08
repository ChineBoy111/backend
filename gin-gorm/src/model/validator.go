package model

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strings"
)

// RegisterCustomValidator 注册自定义模型验证器
// ! 使用自定义模型验证器
//
//	type ModelName struct {
//	    FieldName string `json:"name" binding:"non-admin"`
//	}
//
// ! 类型断言 typeX, ok := x.(Type); ok 表示类型断言是否成功
func RegisterCustomValidator() {
	if customValidator, ok /* 类型断言是否成功 */ := binding.Validator.Engine().(*validator.Validate); ok {
		customValidator.RegisterValidation("non-admin", func(fieldLevel validator.FieldLevel) bool {
			if fieldName, ok /* 类型断言是否成功 */ := fieldLevel.Field().Interface().(string); ok {
				// fieldName 不为空，且不以 admin 开头
				if fieldName != "" && strings.Index(fieldName, "admin") != 0 {
					return true
				}
			}
			return false
		})
	}
}

# api 层（等价于 controller 层）

API, Application Program Interface 应用程序接口

## 结构体标签

结构体标签 (Struct Tag) ：结构体字段的元信息，name:value 形式

例 gorm 标签

```go
type User struct {
	//! primarykey  主键
	ID      uint `gorm:"primarykey"`
	//! index       为 users 表的 deleted 字段创建索引
	Deleted bool `gorm:"index"`
	//! size:32     32 字节
	//! not null    非空
	//! column:name users 表中的字段名为 name
	Username string `gorm:"size:32;not null;column:name"`
}
```

例 json 标签

```go
type RespBody struct {
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
```

## HTTP 状态码

| HTTP 状态码 | e.g.                               | 描述       |
| ----------- | ---------------------------------- | ---------- |
| 2xx         | http.StatusOK ==200                | 成功       |
| 3xx         | http.StatusMovedPermanently == 301 | 重定向     |
| 4xx         | http.StatusNotFound == 404         | 客户端错误 |
| 5xx         | http.StatusBadGateway == 502       | 服务器错误 |

## 判断 x, y 是否深度相等

```go
package reflect

func DeepEqual(x, y interface{}) bool // 通过反射，判断 x, y 是否深度相等
```

## 自定义字段校验器

数据传输对象 [UserLoginDto](../model/user.go)

```go
type UserLoginDto struct {
	//* json:"name"
	//! name      json 中的字段名为 name
	//* binding:"required,not_admin"
	//! required  必填字段，绑定时如果 name 为空则报错
	//! not_admin     自定义字段校验器 ../../model/validator.go
	Username string `json:"username" binding:"required,not_admin" msg:"用户名错误" required_err:"用户名不能为空" not_admin_err:"用户名非法"`
	Password string `json:"password" binding:"required" msg:"密码错误"`
}
```

自定义字段校验器 [validator](./validator.go)

```go

// RegisterCustomValidator 注册自定义字段校验器
// ! 使用自定义字段校验器
//
//	type StructName struct {
//	    FieldName string `json:"name" binding:"not_admin"`
//	}
//
// ! 类型断言 typeX, ok := x.(Type); ok 表示类型断言是否成功
func RegisterCustomValidator() {
	if customValidator, ok /* 类型断言是否成功 */ := binding.Validator.Engine().(*validator.Validate); ok {
		customValidator.RegisterValidation("not_admin" /* 结构体标签名 */, func(fieldLevel validator.FieldLevel) bool {
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
```

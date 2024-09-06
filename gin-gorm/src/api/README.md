## 结构体标签

结构体标签 (Struct Tag) ：结构体字段的元信息

例 gorm

```go
type User struct {
	ID      uint `gorm:"primarykey"` // 主键
	Deleted bool `gorm:"index"`      // 为 users 表的 deleted 字段创建索引
	// 32 字节、非空、users 表中的字段名为 name
	Username string `gorm:"size:32;not null;column:name"`
}
```

例 json

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
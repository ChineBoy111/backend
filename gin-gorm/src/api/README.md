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

| HTTP 状态码 | 描述          |
| ----------- | ------------- |
| 2xx         | http.StatusOK |

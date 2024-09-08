package dto

// DAO: Data Access Object 数据访问对象
// DTO: Data Transfer Object 数据传输对象

type UserLoginDto struct {
	//* json:"name"
	//! name      json 中的字段名为 name
	//* binding:"required,non-admin"
	//! required  必填字段，绑定时如果 name 为空则报错
	//! non-admin 自定义模型验证器 ../../model/validator.go
	Name     string `json:"name" binding:"required,non-admin"`
	Password string `json:"password" binding:"required"`
}

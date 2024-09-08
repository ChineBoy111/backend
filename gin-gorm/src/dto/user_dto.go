package dto

//! ==================== DTO, Data Transfer Object 数据传输对象 ====================

type UserLoginDto struct {
	//* json:"name"
	//! name      json 中的字段名为 name
	//* binding:"required,not_admin"
	//! required  必填字段，绑定时如果 name 为空则报错
	//! not_admin 自定义字段校验器 ../../model/validator.go
	Username string `json:"username" binding:"required,not_admin" msg:"用户名错误" required_err:"用户名不能为空" not_admin_err:"用户名非法"`
	Password string `json:"password" binding:"required" msg:"密码错误"`
}

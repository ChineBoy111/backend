package dto

import (
	"bronya.com/gin-gorm/src/model"
)

//! ==================== DTO, Data Transfer Object 数据传输对象 ====================

type UserLoginDto struct {
	//* json:"username"
	//! username  - json 中的字段名为 username
	//* binding:"required,not_admin"
	//! required  - 必填字段，绑定时如果 name 为空则报错
	//! not_admin - 自定义字段校验器 ../../model/validator.go
	Username string `json:"username" binding:"required,not_admin" msg:"username error" required_err:"username can't be empty" not_admin_err:"username can't be admin"`
	Password string `json:"password" binding:"required" msg:"password error"`
}

type UserInsertDto struct {
	ID uint //! 接收 gorm 生成的主键 ID
	//* form:"username"
	//! username - HTML 表单中，input 标签的 id 为 username
	Username string `json:"username" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Avatar   string
}

func (userInsertDto *UserInsertDto) ToUser() model.User {
	var user model.User
	user.Username = userInsertDto.Username
	user.Password = userInsertDto.Password
	user.Name = userInsertDto.Name
	user.Phone = userInsertDto.Phone
	user.Email = userInsertDto.Email
	return user
}

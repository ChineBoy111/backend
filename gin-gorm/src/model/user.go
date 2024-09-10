package model

import "gorm.io/gorm"

// ! ==================== Model 数据模型 ====================

type User struct {
	//* ID
	//* CreatedAt 创建时间
	//* UpdatedAt 更新时间
	//* DeletedAt 删除时间、是否删除
	gorm.Model
	Username string `gorm:"size:64;not null"`  // 用户名，非空
	Password string `gorm:"size:255;not null"` // 密码，非空
	Name     string `gorm:"size:128"`          // 姓名
	Phone    string `gorm:"size:11"`           // 电话
	Email    string `gorm:"size:128"`          // 邮箱
	Avatar   string `gorm:"size:255"`          // 头像
}

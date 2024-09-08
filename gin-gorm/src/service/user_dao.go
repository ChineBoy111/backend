package service

import "gorm.io/gorm"

//! DAO Data Access Object

var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{}
	}
	return userDao
}

type UserDao struct {
	db *gorm.DB
}

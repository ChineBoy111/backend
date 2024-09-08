package dao

import (
	"bronya.com/gin-gorm/src/data"
	"bronya.com/gin-gorm/src/global"
	"gorm.io/gorm"
)

// ! ==================== DAO, Data Access Object 数据访问对象 ====================

type UserDao struct {
	database *gorm.DB
}

// ! UserDao 单例
var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			database: global.Database,
		}
	}
	return userDao
}

func (userDao *UserDao) SelectUserByUsernameAndPassword(username, password string) data.User {
	var user data.User
	userDao.database.Where("username = ? and password = ?", username, password).First(&user)
	return user
}

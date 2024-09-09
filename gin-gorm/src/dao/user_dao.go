package dao

import (
	"bronya.com/gin-gorm/src/dto"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/model"
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

func (userDao *UserDao) SelectUserByUsernameAndPassword(username, password string) model.User {
	var user model.User
	userDao.database.Model(&model.User{}).Where("username = ? and password = ?", username, password).First(&user)
	return user
}

func (userDao *UserDao) InsertUser(userInsertDto *dto.UserInsertDto) error {
	user := userInsertDto.ToUser()
	err := userDao.database.Model(&model.User{}).Save(&user).Error
	if err == nil {
		userInsertDto.ID = user.ID //! 接收 gorm 生成的主键 ID
	}
	return err
}

func (userDao *UserDao) SelectUserByUsername(username string) model.User {
	var user model.User
	userDao.database.Model(&model.User{}).Where("username = ?", username).First(&user)
	return user
}

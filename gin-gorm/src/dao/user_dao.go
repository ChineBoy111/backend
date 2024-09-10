package dao

import (
	"bronya.com/gin-gorm/src/dto"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/model"
	"errors"
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

func (userDao *UserDao) InsertUser(userInsertDto *dto.UserInsertDto) error {
	user := userInsertDto.ToUser()
	err := userDao.database.Model(&model.User{}).Save(&user).Error
	if err == nil {
		userInsertDto.ID = user.ID //! 接收 gorm 生成的主键 ID
	}
	return err
}

func (userDao *UserDao) SelectUserByUsernameAndPassword(username, password string) (model.User, error) {
	var user model.User
	err := userDao.database.Model(&model.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 || password != user.Password {
		return user, errors.New("username or password error")
	}
	return user, nil
}

func (userDao *UserDao) SelectUserByUsername(username string) (model.User, error) {
	var user model.User
	err := userDao.database.Model(&model.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("username error")
	}
	return user, nil
}

func (userDao *UserDao) SelectUserById(id uint) (model.User, error) {
	var user model.User
	err := userDao.database.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

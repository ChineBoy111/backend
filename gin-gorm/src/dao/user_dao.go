package dao

import (
	"bronya.com/gin-gorm/src/data"
	"bronya.com/gin-gorm/src/dto"
	"bronya.com/gin-gorm/src/global"
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
	err := userDao.database.Model(&data.User{}).Save(&user).Error
	if err == nil {
		userInsertDto.ID = user.ID //! 接收 gorm 生成的主键 ID
	}
	return err
}

func (userDao *UserDao) SelectUserByUsernameAndPassword(username, password string) (data.User, error) {
	var user data.User
	err := userDao.database.Model(&data.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 || password != user.Password {
		return user, errors.New("username or password error")
	}
	return user, nil
}

func (userDao *UserDao) SelectUserByUsername(username string) (data.User, error) {
	var user data.User
	err := userDao.database.Model(&data.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("username error")
	}
	return user, nil
}

func (userDao *UserDao) SelectUserById(id uint) (data.User, error) {
	var user data.User
	err := userDao.database.Model(&data.User{}).First(&user, id).Error
	return user, err
}

func (userDao *UserDao) SelectUserByPage(paginateDto *dto.PaginateDto) ([]data.User, int64, error) {
	var userArr []data.User //! 接收分页查询结果
	var total int64         //! 接收总记录数
	err := userDao.database.Model(&data.User{}).
		Scopes(GetPaginateFunc(paginateDto)). //! 传递分页函数
		Find(&userArr).                       //! 获取分页查询结果
		Offset(-1).Limit(-1).                 //! 取消分页查询条件
		Count(&total).                        //! 获取总记录数
		Error                                 //! 获取可能的错误
	return userArr, total, err
}

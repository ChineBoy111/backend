package service

import (
	"bronya.com/gin-gorm/src/dao"
	"bronya.com/gin-gorm/src/dao/data"
	"bronya.com/gin-gorm/src/service/dto"
	"errors"
)

type UserService struct {
	UserDao *dao.UserDao
}

// ! UserService 单例
var userService *UserService

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			UserDao: dao.NewUserDao(), //! 组合 UserDao
		}
	}
	return userService
}
func (userService UserService) Login(userLoginDto dto.UserLoginDto) (data.User, error) {
	var err error
	user := userService.UserDao.SelectUserByUsernameAndPassword(userLoginDto.Username, userLoginDto.Password)
	if user.ID == 0 {
		err = errors.New("username or password error")
	}
	return user, err
}

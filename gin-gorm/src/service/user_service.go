package service

import (
	"bronya.com/gin-gorm/src/dao"
	"bronya.com/gin-gorm/src/dto"
	"bronya.com/gin-gorm/src/model"
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
func (userService *UserService) Login(userLoginDto *dto.UserLoginDto) (model.User, error) {
	var err error
	user := userService.UserDao.SelectUserByUsernameAndPassword(userLoginDto.Username, userLoginDto.Password)
	if user.ID == 0 {
		err = errors.New("username or password error")
	}
	return user, err
}

func (userService *UserService) AddUser(userInsertDto *dto.UserInsertDto) error {
	if userService.UserDao.SelectUserByUsername(userInsertDto.Username).ID != 0 {
		return errors.New("username already exists")
	}
	return userService.UserDao.InsertUser(userInsertDto)
}

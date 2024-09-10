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
	user, err := userService.UserDao.SelectUserByUsernameAndPassword(userLoginDto.Username, userLoginDto.Password)
	return user, err
}

func (userService *UserService) InsertUser(userInsertDto *dto.UserInsertDto) error {
	if _, err := userService.UserDao.SelectUserByUsername(userInsertDto.Username); err == nil {
		return errors.New("username exists")
	}
	return userService.UserDao.InsertUser(userInsertDto)
}

func (userService *UserService) SelectUserById(commonIdDto *dto.CommonIdDto) (model.User, error) {
	return userService.UserDao.SelectUserById(commonIdDto.ID)
}

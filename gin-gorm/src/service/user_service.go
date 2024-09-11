package service

import (
	"bronya.com/gin-gorm/src/dao"
	"bronya.com/gin-gorm/src/data"
	"bronya.com/gin-gorm/src/dto"
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

func (userService *UserService) SelectUser(userLoginDto *dto.UserLoginDto) (data.User, error) {
	var err error
	user, err := userService.UserDao.SelectUser(userLoginDto)
	return user, err
}

func (userService *UserService) InsertUser(userInsertDto *dto.UserInsertDto) error {
	if _, err := userService.UserDao.SelectUserByUsername(userInsertDto.Username); err == nil {
		return errors.New("username exists")
	}
	return userService.UserDao.InsertUser(userInsertDto)
}

func (userService *UserService) SelectUserById(idDto *dto.IdDto) (data.User, error) {
	return userService.UserDao.SelectUserById(idDto.Id)
}

func (userService *UserService) SelectPaginatedUsers(paginateDto *dto.PaginateDto) ([]data.User, int64, error) {
	return userService.UserDao.SelectPaginatedUsers(paginateDto)
}

func (userService *UserService) UpdateUser(userUpdateDto *dto.UserUpdateDto) error {
	if userUpdateDto.Id <= 0 {
		return errors.New("Id error")
	}
	return userService.UserDao.UpdateUser(userUpdateDto)
}

func (userService *UserService) DeleteUserById(idDto *dto.IdDto) (data.User, error) {
	return userService.UserDao.DeleteUserById(idDto.Id)
}

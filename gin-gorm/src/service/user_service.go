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

func (userService *UserService) SelectUserByUsernameAndPassword(userLoginDto *dto.UserLoginDto) (data.User, error) {
	var err error
	user, err := userService.UserDao.SelectUserByUsernameAndPassword(userLoginDto)
	return user, err
}

func (userService *UserService) InsertUser(userInsertDto *dto.UserInsertDto) error {
	if _, err := userService.UserDao.SelectUserByUsername(userInsertDto.Username); err == nil {
		return errors.New("username exists")
	}
	return userService.UserDao.InsertUser(userInsertDto)
}

func (userService *UserService) SelectUserById(commonIdDto *dto.IdDto) (data.User, error) {
	return userService.UserDao.SelectUserById(commonIdDto.ID)
}

func (userService *UserService) SelectPaginatedUser(paginateDto *dto.PaginateDto) ([]data.User, int64, error) {
	return userService.UserDao.SelectPaginatedUser(paginateDto)
}

func (userService *UserService) UpdateUser(userUpdateDto *dto.UserUpdateDto) error {
	if userUpdateDto.ID <= 0 {
		return errors.New("ID error")
	}
	return userService.UserDao.UpdateUser(userUpdateDto)
}

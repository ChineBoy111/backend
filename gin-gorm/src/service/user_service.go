package service

import (
	"bronya.com/gin-gorm/src/dao"
	"bronya.com/gin-gorm/src/data"
	"bronya.com/gin-gorm/src/dto"
	"bronya.com/gin-gorm/src/global"
	"bronya.com/gin-gorm/src/util"
	"context"
	"errors"
	"github.com/spf13/viper"
	"strconv"
	"time"
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

func (userService *UserService) Login(loginDto *dto.LoginDto) (data.User, string, error) {
	var err error
	var token string
	user, err := userService.UserDao.SelectUserByUsername(loginDto.Username)
	// 登录失败
	if err != nil {
		return user, token, err
	}
	// 登录失败
	if !util.IsEquivalent(user.Password /* hashStr */, loginDto.Password /* password */) {
		err = errors.New("username or password error")
		return user, token, err
	}
	token, err = util.GenToken(user.ID, user.Username)
	// 登录失败
	if err != nil {
		return user, token, err
	}
	// 登录成功
	//! 使用 redis 设置 token 的过期时间
	expire := viper.GetDuration("redis.expire")
	global.RedisCli.Set(context.Background(), "user"+strconv.Itoa(int(user.ID)), token, expire*time.Second)
	return user, token, err
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
		return errors.New("id error")
	}
	return userService.UserDao.UpdateUser(userUpdateDto)
}

func (userService *UserService) DeleteUserById(idDto *dto.IdDto) (data.User, error) {
	return userService.UserDao.DeleteUserById(idDto.Id)
}

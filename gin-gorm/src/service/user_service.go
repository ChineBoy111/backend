package service

import "bronya.com/gin-gorm/src/data"

func (userDao *UserDao) GetUserByUsername(username, password string) data.User {
	var user data.User
	userDao.db.Where("username = ? and password = ?", username, password).First(&user)
	return user
}

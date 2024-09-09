package conf

import (
	"bronya.com/gin-gorm/src/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnMysql() (*gorm.DB, error) {

	gormLogMod := gormLogger.Info

	if viper.GetString("build.type") == "Release" {
		gormLogMod = gormLogger.Error
	}

	//! 开启会话 session
	dbSession, err := gorm.Open(mysql.Open(viper.GetString("db.mysql.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, //! 单数表名，默认 false
			TablePrefix:   "t_",  //! 表名前缀 t_users
		},
		Logger: gormLogger.Default.LogMode(gormLogMod),
	})

	if err != nil {
		return nil, err
	}

	db, _ := dbSession.DB()
	//* 最大空闲连接数，默认 2
	db.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	//* 最大连接数，默认 0，表示无限制
	db.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	//* 最长连接时间，默认 0，表示无限制
	//! `type Duration int64`
	//! 使用 type.Duration() 强制类型转换
	db.SetConnMaxLifetime(viper.GetDuration("db.connMaxLifetime"))

	//! 从 go 结构体自动迁移到数据库表，创建表
	dbSession.AutoMigrate(&model.User{}) //* 传递指向一个 model.User 对象的指针
	return dbSession, nil
}

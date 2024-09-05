package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//  全局变量

var (
	Logger *zap.SugaredLogger
	DBSession *gorm.DB
)

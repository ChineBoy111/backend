package api

import (
	"bronya.com/gin-gorm/src/service"
	"github.com/gin-gonic/gin"
)

type HostApi struct {
	HostService *service.HostService
}

// ! HostApi 单例
var hostApi *HostApi

func NewHostApi() *HostApi {
	if hostApi == nil {
		hostApi = &HostApi{
			HostService: service.NewHostService(),
		}
	}
	return hostApi
}

func (hostApi HostApi) Shutdown(c *gin.Context) { //! 不使用指针接收

}

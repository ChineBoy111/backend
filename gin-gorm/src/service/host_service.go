package service

import "bronya.com/gin-gorm/src/service/dto"

type HostService struct {
}

// ! HostService 单例
var hostService *HostService

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (hostService *HostService) Shutdown(hostShutdownDto dto.HostShutdownDto) {
	hostIp := hostShutdownDto.HostIp

}

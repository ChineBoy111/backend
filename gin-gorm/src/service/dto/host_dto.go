package dto

type HostShutdownDto struct {
	HostIp string `json:"host_ip" binding:"required" msg:"主机 IP 错误"`
}

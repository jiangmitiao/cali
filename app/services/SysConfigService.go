package services

import "github.com/jiangmitiao/cali/app/models"

type SysConfigService struct {
}

func (service SysConfigService) Get(key string) string {
	sysConfig := models.SysConfig{}
	localEngine.Where("key = ?", key).Get(&sysConfig)
	return sysConfig.Value
}

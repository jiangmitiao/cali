package services

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
)

type SysConfigService struct {
}

func (service SysConfigService) Get(key string) string {
	sysConfig := models.SysConfig{}
	localEngine.Where("key = ?", key).Get(&sysConfig)
	return sysConfig.Value
}

func (service SysConfigService) QuerySysConfigs(limit, start int) []models.SysConfig {
	sysConfigs := make([]models.SysConfig, 0)
	localEngine.Limit(limit, start).Find(&sysConfigs)
	return sysConfigs
}

func (service SysConfigService) UpdateConfig(sysConfig models.SysConfig) bool {
	_, err := localEngine.Id(sysConfig.Id).Cols("key", "value").Update(sysConfig)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (service SysConfigService) AddSysConfig(sysConfig models.SysConfig) bool {
	sysConfig.Id = uuid.New().String()
	_, err := localEngine.Insert(sysConfig)
	if err == nil {
		return true
	} else {
		return false
	}
}

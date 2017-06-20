package services

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
)

type SysConfigService struct {
}

func (service SysConfigService) Get(key string) models.SysConfig {
	sysConfig := models.SysConfig{}
	localEngine.Where("key = ?", key).Get(&sysConfig)
	return sysConfig
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
	if count, err := localEngine.Where("key = ?", sysConfig.Key).Count(models.SysConfig{}); err == nil {
		if count == 1 {
			if _, err := localEngine.Where("key = ?", sysConfig.Key).Cols("value").Update(sysConfig); err == nil {
				return true
			} else {
				return false
			}
		} else {
			sysConfig.Id = uuid.New().String()
			if _, err := localEngine.InsertOne(sysConfig); err == nil {
				return true
			} else {
				return false
			}
		}
	} else {
		return false
	}
}

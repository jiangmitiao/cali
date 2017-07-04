package services

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"time"
)

type SysConfigService struct {
}

func (service SysConfigService) Get(key string) models.SysConfig {
	sysConfig := models.SysConfig{}
	engine.Where("ikey like ?", key).Get(&sysConfig)
	return sysConfig
}

func (service SysConfigService) QuerySysConfigs(limit, start int) []models.SysConfig {
	sysConfigs := make([]models.SysConfig, 0)
	engine.Limit(limit, start).Find(&sysConfigs)
	return sysConfigs
}

func (service SysConfigService) UpdateConfig(sysConfig models.SysConfig) bool {
	sysConfig.UpdatedAt = time.Now().Unix()
	_, err := engine.Id(sysConfig.Id).Cols("ikey", "value", "updated").Update(sysConfig)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (service SysConfigService) AddSysConfig(sysConfig models.SysConfig) bool {
	if count, err := engine.Where("ikey = ?", sysConfig.Ikey).Count(models.SysConfig{}); err == nil {
		if count == 1 {
			sysConfig.UpdatedAt = time.Now().Unix()
			if _, err := engine.Where("ikey = ?", sysConfig.Ikey).Cols("value", "updated").Update(sysConfig); err == nil {
				return true
			} else {
				return false
			}
		} else {
			sysConfig.Id = uuid.New().String()
			sysConfig.UpdatedAt = time.Now().Unix()
			sysConfig.CreatedAt = time.Now().Unix()
			if _, err := engine.InsertOne(sysConfig); err == nil {
				return true
			} else {
				return false
			}
		}
	} else {
		return false
	}
}

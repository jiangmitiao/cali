package services

import (
	"sync"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/google/uuid"
)

type SysStatusService struct {
	lock *sync.Mutex
}

func (service SysStatusService) Get(key string) string {
	sysStatus := models.SysStatus{}
	localEngine.Where("key = ?", key).Get(&sysStatus)
	return sysStatus.Value
}

func (service SysStatusService) QuerySysStatus(limit, start int) []models.SysStatus {
	sysStatus := make([]models.SysStatus, 0)
	localEngine.Limit(limit, start).Find(&sysStatus)
	return sysStatus
}

func (service SysStatusService) UpdateStatus(sysStatus models.SysStatus) bool {
	_, err := localEngine.Id(sysStatus.Id).Cols("key", "value").Update(sysStatus)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (service SysStatusService) AddSysStatus(sysStatus models.SysStatus) bool {
	sysStatus.Id = uuid.New().String()
	_, err := localEngine.Insert(sysStatus)
	if err == nil {
		return true
	} else {
		return false
	}
}

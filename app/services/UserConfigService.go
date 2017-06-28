package services

import "github.com/jiangmitiao/cali/app/models"

type UserConfigService struct {
}

func (service UserConfigService) GetUserConfig(userId string) (models.UserConfig, bool) {
	//if not exist ,then insert a new
	config := models.UserConfig{}
	if ok, _ := engine.Where("user_info = ?", userId).Get(&config); ok {
		return config, ok
	} else {
		config = models.NewUserConfig(userId)
		if _, err := engine.InsertOne(config); err == nil {
			return config, true
		} else {
			return config, false
		}
	}
}

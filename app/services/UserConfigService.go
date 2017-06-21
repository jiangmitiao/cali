package services

import "github.com/jiangmitiao/cali/app/models"

type UserConfigService struct {
}

func (service UserConfigService) GetUserConfig(userId string) (models.UserConfig, bool) {
	//if not exist ,then insert a new
	config := models.UserConfig{}
	if ok, _ := localEngine.Where("user_info = ?", userId).Get(&config); ok {
		return config, ok
	} else {
		config = models.NewUserConfig(userId)
		if _, err := localEngine.InsertOne(config); err == nil {
			return config, true
		} else {
			return config, false
		}
	}
}

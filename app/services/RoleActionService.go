package services

import "github.com/jiangmitiao/cali/app/models"

type RoleActionService struct {
}

func (service RoleActionService) GetRoleActionByControllerMethodRole(controller, method, role string) models.RoleAction {
	roleAction := models.RoleAction{}
	localEngine.Where("controller = ?", controller).Where("method = ?", method).Where("role = ?", role).Get(&roleAction)
	return roleAction
}

package services

import "github.com/jiangmitiao/cali/app/models"

type UserRoleService struct {

}

func (service UserRoleService) GetRoleByUser(userId string)models.Role  {
	role := models.Role{}
	localEngine.SQL("select role.* from role ,user_info_role_link where role.id=user_info_role_link.role and user_info_role_link.user_info="+userId).
		Get(&role)
	return role
}

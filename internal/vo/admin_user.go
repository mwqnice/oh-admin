package vo

import "github.com/mwqnice/oh-admin/internal/model"

//AdminUserInfoVo 用户信息Vo
type AdminUserInfoVo struct {
	*model.AdminUser
	GenderName string      `json:"gender_name"` // 性别
	RoleList   interface{} `json:"role_list"`   // 角色列表
	RoleIds    []int       `json:"role_ids"`    // 角色ID
}

package dto

//GetAdminRoleListRequest 获取角色列表
type GetAdminRoleListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   string `form:"name" json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

//CreateAdminRoleRequest 添加角色
type CreateAdminRoleRequest struct {
	Name   string `form:"name" binding:"required" json:"name"`
	Code   string `form:"code" binding:"required" json:"code"`
	Status string `form:"status" binding:"required" json:"status"`
	Sort   string `form:"sort" json:"sort"`
}

//UpdateAdminRoleRequest 更新角色
type UpdateAdminRoleRequest struct {
	Id     string `form:"id" binding:"required" json:"id"`
	Name   string `form:"name" binding:"required" json:"name"`
	Code   string `form:"code" binding:"required" json:"code"`
	Status string `form:"status" binding:"required" json:"status"`
	Sort   string `form:"sort" json:"sort"`
}

//SetAdminRoleStatusRequest 设置状态
type SetAdminRoleStatusRequest struct {
	Id     string `form:"id" binding:"required" json:"id"`
	Status string `form:"status"    binding:"required" json:"status"`
}

//SaveAdminRoleMenuRequest 保存角色菜单
type SaveAdminRoleMenuRequest struct {
	RoleId  int    `form:"role_id" binding:"required" json:"role_id"`
	MenuIds string `form:"menu_ids" binding:"required" json:"menu_ids"`
}

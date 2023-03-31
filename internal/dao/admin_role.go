package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

//GetAdminRoleList 获取角色列表
func (d *Dao) GetAdminRoleList(ctx context.Context, params *dto.GetAdminRoleListRequest) ([]*model.AdminRole, int64, error) {
	adminRole := model.AdminRole{}
	return adminRole.List(ctx, d.engine, params)
}

//CreateAdminRole 创建角色
func (d *Dao) CreateAdminRole(ctx context.Context, role *model.AdminRole) error {
	return role.Create(ctx, d.engine)
}

//GetAdminRoleInfo 获取角色详情
func (d *Dao) GetAdminRoleInfo(ctx context.Context, id int64) (*model.AdminRole, error) {
	adminRole := model.AdminRole{}
	return adminRole.Get(ctx, d.engine, id)
}

//UpdateAdminRole 更新角色
func (d *Dao) UpdateAdminRole(ctx context.Context, role *model.AdminRole) error {
	return role.Update(ctx, d.engine)
}

//DeleteAdminRole 删除角色
func (d *Dao) DeleteAdminRole(ctx context.Context, ids []int64) error {
	adminRole := model.AdminRole{}
	return adminRole.DeleteBatch(ctx, d.engine, ids)
}

//GetAdminRoleMenuList 获取角色菜单列表
func (d *Dao) GetAdminRoleMenuList(ctx context.Context, roleMenu *model.AdminRoleMenu) ([]*model.AdminRoleMenu, error) {
	return roleMenu.List(ctx, d.engine)
}

//DeleteAdminRoleMenu 删除角色菜单列表
func (d *Dao) DeleteAdminRoleMenu(ctx context.Context, roleMenu *model.AdminRoleMenu) error {
	return roleMenu.Delete(ctx, d.engine)
}

//BatchCreateAdminRoleMenu 批量创建角色菜单列表
func (d *Dao) BatchCreateAdminRoleMenu(ctx context.Context, record []*model.AdminRoleMenu) (int64, error) {
	roleMenu := model.AdminRoleMenu{}
	return roleMenu.BatchCreate(ctx, d.engine, record)
}

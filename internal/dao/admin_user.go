package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

//GetAdminUser 获取用户信息
func (d *Dao) GetAdminUser(ctx context.Context, adminUser *model.AdminUser) (*model.AdminUser, error) {
	return adminUser.Get(ctx, d.engine)
}

//UpdateAdminUser 更新用户
func (d *Dao) UpdateAdminUser(ctx context.Context, adminUser *model.AdminUser) error {
	return adminUser.Update(ctx, d.engine)
}

//GetAdminUserList 获取管理员列表
func (d *Dao) GetAdminUserList(ctx context.Context, params *dto.GetAdminUserListRequest) ([]*model.AdminUser, int64, error) {
	adminUser := model.AdminUser{}
	return adminUser.List(ctx, d.engine, params)
}

//CreateAdminUser 创建用户
func (d *Dao) CreateAdminUser(ctx context.Context, adminUser *model.AdminUser) (*model.AdminUser, error) {
	return adminUser.Create(ctx, d.engine)
}

//GetAdminUserRoleList 获取用户角色列表
func (d *Dao) GetAdminUserRoleList(ctx context.Context, userId int) ([]*model.AdminRole, error) {
	adminUserRole := model.AdminUserRole{}
	var list []*model.AdminRole
	err := d.engine.WithContext(ctx).Table(adminUserRole.TableName()+" ur").
		Joins("left join oh_admin_role as role on ur.role_id = role.id").
		Order("role.id asc").
		Where("ur.user_id=? AND role.`status`=1", userId).
		Select("role.*").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

//BatchCreateAdminUserRole 批量创建用户角色
func (d *Dao) BatchCreateAdminUserRole(ctx context.Context, record []*model.AdminUserRole) (int64, error) {
	userRole := model.AdminUserRole{}
	return userRole.BatchCreate(ctx, d.engine, record)
}

//DeleteAdminUserRole 删除用户角色
func (d *Dao) DeleteAdminUserRole(ctx context.Context, userRole *model.AdminUserRole) error {
	return userRole.Delete(ctx, d.engine)
}

//DeleteAdminUser 删除用户
func (d *Dao) DeleteAdminUser(ctx context.Context, id int64) error {
	adminUser := model.AdminUser{}
	return adminUser.Delete(ctx, d.engine, int(id))
}

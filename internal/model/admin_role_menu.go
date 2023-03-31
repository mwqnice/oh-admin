package model

import (
	"context"
	"github.com/mwqnice/oh-admin/global"
	"gorm.io/gorm"
)

type AdminRoleMenu struct {
	*Model
	RoleId int `gorm:"column:role_id;default:0;NOT NULL" json:"role_id"` // 角色id
	MenuId int `gorm:"column:menu_id;default:0;NOT NULL" json:"menu_id"` // 菜单id
}

func (roleMenu *AdminRoleMenu) TableName() string {
	return global.DatabaseSetting.TablePrefix + "admin_role_menu"
}

//List 列表
func (roleMenu *AdminRoleMenu) List(ctx context.Context, db *gorm.DB) ([]*AdminRoleMenu, error) {

	var list []*AdminRoleMenu
	db = db.WithContext(ctx).Table(roleMenu.TableName())
	if roleMenu.RoleId != 0 {
		db = db.Where("role_id = ?", roleMenu.RoleId)
	}
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

//Delete 删除
func (roleMenu *AdminRoleMenu) Delete(ctx context.Context, db *gorm.DB) error {
	db = db.WithContext(ctx)
	if roleMenu.RoleId != 0 {
		db = db.Where("role_id = ?", roleMenu.RoleId)
	}
	return db.Delete(&roleMenu).Error
}

//BatchCreate 批量创建
func (roleMenu *AdminRoleMenu) BatchCreate(ctx context.Context, db *gorm.DB, record []*AdminRoleMenu) (int64, error) {
	result := db.WithContext(ctx).Create(&record)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

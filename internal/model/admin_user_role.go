package model

import (
	"context"
	"github.com/mwqnice/oh-admin/global"
	"gorm.io/gorm"
)

type AdminUserRole struct {
	*Model
	UserId int `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"` // 用户id
	RoleId int `gorm:"column:role_id;default:0;NOT NULL" json:"role_id"` // 角色id
}

func (userRole *AdminUserRole) TableName() string {
	return global.DatabaseSetting.TablePrefix + "admin_user_role"
}

//BatchCreate 批量创建
func (userRole *AdminUserRole) BatchCreate(ctx context.Context, db *gorm.DB, record []*AdminUserRole) (int64, error) {
	result := db.WithContext(ctx).Create(&record)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

//Delete 删除
func (userRole *AdminUserRole) Delete(ctx context.Context, db *gorm.DB) error {
	db = db.WithContext(ctx)
	if userRole.UserId != 0 {
		db = db.Where("user_id = ?", userRole.UserId)
	}
	return db.Delete(&userRole).Error
}

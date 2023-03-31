package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"gorm.io/gorm"
)

type AdminRole struct {
	*Model
	Name   string `gorm:"column:name;NOT NULL" json:"name"`      // 角色名称
	Code   string `gorm:"column:code;NOT NULL" json:"code"`      // 角色编码
	Sort   int    `gorm:"column:sort;default:0" json:"sort"`     // 排序
	Status int    `gorm:"column:status;default:1" json:"status"` // 状态 1-启用2-禁用
}

func (role *AdminRole) TableName() string {
	return global.DatabaseSetting.TablePrefix + "admin_role"
}

//List 列表
func (m *AdminRole) List(ctx context.Context, db *gorm.DB, params *dto.GetAdminRoleListRequest) ([]*AdminRole, int64, error) {
	db = db.WithContext(ctx).Table(m.TableName())
	if params.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	if params.Status != "" {
		db = db.Where("status = ?", convert.Int(params.Status))
	}
	var count int64
	db.Count(&count)
	var list []*AdminRole
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id asc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

//Create 创建
func (m *AdminRole) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&m).Error
}

//Get 查询
func (m *AdminRole) Get(ctx context.Context, db *gorm.DB, id int64) (*AdminRole, error) {
	if err := db.WithContext(ctx).Where("id = ? ", id).First(&m).Error; err != nil && err != gorm.ErrRecordNotFound {
		return m, err
	}
	return m, nil
}

//Update 更新
func (m *AdminRole) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&m).Error
}

//DeleteBatch 批量删除
func (m *AdminRole) DeleteBatch(ctx context.Context, db *gorm.DB, ids []int64) error {
	return db.WithContext(ctx).Where("id in ?", ids).Delete(&m).Error
}

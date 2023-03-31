package model

import (
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AdminMenu struct {
	*Model
	Name       string `gorm:"column:name;NOT NULL" json:"name"`           // 菜单名称
	Icon       string `gorm:"column:icon" json:"icon"`                    // 图标
	Url        string `gorm:"column:url" json:"url"`                      // URL地址
	Pid        int    `gorm:"column:pid;default:0;NOT NULL" json:"pid"`   // 上级ID
	Type       int    `gorm:"column:type;default:0;NOT NULL" json:"type"` // 类型：1模块 2导航 3菜单 4节点
	IsShow     int    `gorm:"column:is_show;default:1" json:"is_show"`    // 是否显示：1显示 2不显示
	Sort       int    `gorm:"column:sort" json:"sort"`                    // 显示顺序
	Permission string `gorm:"column:permission" json:"permission"`        // 权限标识
	Remark     string `gorm:"column:remark" json:"remark"`                // 菜单备注
	Status     int    `gorm:"column:status;default:1" json:"status"`      // 状态1-在用 2-禁用
}

func (m *AdminMenu) TableName() string {
	return global.DatabaseSetting.TablePrefix + "admin_menu"
}

func (m *AdminMenu) List(ctx context.Context, db *gorm.DB, params *dto.GetAdminMenuListRequest) ([]*AdminMenu, int64, error) {
	db = db.WithContext(ctx).Table(m.TableName())
	if params.Status > 0 {
		db = db.Where("status = ?", params.Status)
	}

	var count int64
	db.Count(&count)
	var list []*AdminMenu
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("sort asc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

//Get 根据条件查询单条数据
func (m *AdminMenu) Get(ctx context.Context, db *gorm.DB) (*AdminMenu, error) {
	db = db.WithContext(ctx).Table(m.TableName())
	if m.Permission != "" {
		db = db.Where("permission = ? ", m.Permission)
	}
	if m.Model != nil && m.ID != 0 {
		db = db.Where("id = ? ", m.ID)
	}
	var adminMenu *AdminMenu
	if err := db.First(&adminMenu).Error; err != nil && err != gorm.ErrRecordNotFound {
		return adminMenu, err
	}
	return adminMenu, nil
}

//Create 插入数据
func (m *AdminMenu) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&m).Error
}

//Update 更新
func (m *AdminMenu) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&m).Error
}

//Delete 删除
func (m *AdminMenu) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&m).Error
}

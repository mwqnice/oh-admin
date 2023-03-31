package model

import (
	"fmt"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Link struct {
	*Model
	Name   string `gorm:"column:name;NOT NULL" json:"name"`               // 友链名称
	Url    string `gorm:"column:url;NOT NULL" json:"url"`                 // 友链地址
	Image  string `gorm:"column:image" json:"image"`                      // logo
	Status int    `gorm:"column:status;default:1;NOT NULL" json:"status"` // 状态 1-启用 2-禁用
	Sort   int    `gorm:"column:sort;default:0;NOT NULL" json:"sort"`     // 排序
}

func (l *Link) TableName() string {
	return global.DatabaseSetting.TablePrefix + "link"
}

func (l *Link) List(ctx context.Context, db *gorm.DB, params *dto.GetLinkListRequest) ([]*Link, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Status > 0 {
		db = db.Where("status = ?", params.Status)
	}
	if params.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*Link
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("sort asc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

//Get 根据条件查询单条数据
func (l *Link) Get(ctx context.Context, db *gorm.DB) (*Link, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.Model != nil && l.ID != 0 {
		db = db.Where("id = ? ", l.ID)
	}
	var link *Link
	if err := db.First(&link).Error; err != nil && err != gorm.ErrRecordNotFound {
		return link, err
	}
	return link, nil
}

//Create 插入数据
func (l *Link) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&l).Error
}

//Update 更新
func (l *Link) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&l).Error
}

//Delete 删除
func (l *Link) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&l).Error
}

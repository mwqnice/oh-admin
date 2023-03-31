package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type AdminUser struct {
	*Model
	Realname  string `gorm:"column:realname;NOT NULL" json:"realname"`               // 真实姓名
	Username  string `gorm:"column:username;NOT NULL" json:"username"`               // 登录用户名
	Gender    int    `gorm:"column:gender;default:0;NOT NULL" json:"gender"`         // 性别:1男 2女 3保密
	Avatar    string `gorm:"column:avatar;NOT NULL" json:"avatar"`                   // 头像
	Mobile    string `gorm:"column:mobile;NOT NULL" json:"mobile"`                   // 手机号
	Email     string `gorm:"column:email;NOT NULL" json:"email"`                     // 邮箱地址
	Address   string `gorm:"column:address;NOT NULL" json:"address"`                 // 地址
	Password  string `gorm:"column:password;NOT NULL" json:"password"`               // 登录密码
	Salt      string `gorm:"column:salt;NOT NULL" json:"salt"`                       // 盐加密
	Intro     string `gorm:"column:intro" json:"intro"`                              // 备注
	Status    int    `gorm:"column:status;default:1;NOT NULL" json:"status"`         // 状态：1正常 2禁用
	LoginNum  int    `gorm:"column:login_num;default:0;NOT NULL" json:"login_num"`   // 登录次数
	LoginIp   string `gorm:"column:login_ip;NOT NULL" json:"login_ip"`               // 最近登录ip
	LoginTime int64  `gorm:"column:login_time;default:0;NOT NULL" json:"login_time"` // 最近登录时间
}

func (a *AdminUser) TableName() string {
	return global.DatabaseSetting.TablePrefix + "admin_user"
}

//Get 获取一条记录
func (a *AdminUser) Get(ctx context.Context, db *gorm.DB) (*AdminUser, error) {
	var admin *AdminUser
	db = db.WithContext(ctx)
	if a.Model != nil && a.ID != 0 {
		db = db.Where("id = ?", a.ID)
	}
	if a.Username != "" {
		db = db.Where("username = ?", a.Username)
	}
	err := db.First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return admin, err
	}
	return admin, nil
}

//Update 更新
func (a *AdminUser) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Model(&a).Updates(&a).Error
}

//Create 创建
func (a *AdminUser) Create(ctx context.Context, db *gorm.DB) (*AdminUser, error) {
	if err := db.WithContext(ctx).Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

//List 列表
func (a *AdminUser) List(ctx context.Context, db *gorm.DB, params *dto.GetAdminUserListRequest) ([]*AdminUser, int64, error) {
	db = db.WithContext(ctx).Table(a.TableName())
	if params.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*AdminUser
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

//Delete 删除
func (m *AdminUser) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&m).Error
}

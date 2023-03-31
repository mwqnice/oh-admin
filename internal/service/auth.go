package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/utils"
	"github.com/mwqnice/oh-admin/pkg/utils/gconv"
	"time"
)

//AdminUserLogin 后台系统登录
func (svc *Service) AdminUserLogin(ctx context.Context, params *dto.AdminUserLoginRequest) (*model.AdminUser, error) {
	// 查询用户
	user, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Username: params.UserName})
	if err != nil && user.Model == nil {
		fmt.Println(1231)
		return nil, errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd, _ := utils.EncodeMD5(params.Password + user.Salt)
	if user.Password != pwd {
		return nil, errors.New("密码不正确")
	}
	// 判断当前用户状态
	if user.Status != model.STATE_OPEN {
		return nil, errors.New("您的账号已被禁用,请联系管理员")
	}

	// 更新登录时间、登录IP
	user.LoginIp = params.Ip
	user.LoginTime = time.Now().Unix()
	user.LoginNum++
	svc.dao.UpdateAdminUser(ctx, user)

	return user, nil
}

//CheckAdminLogin 判断用户登录状态
func (svc *Service) CheckAdminLogin(ctx *gin.Context) bool {
	// 初始化session对象
	session := sessions.Default(ctx)
	// 获取用户ID
	userId := session.Get("adminUserId")
	return userId != nil
}

//AdminLoginOut 退出登录
func (svc *Service) AdminLoginOut(ctx *gin.Context) error {
	// 初始化session对象
	session := sessions.Default(ctx)
	// 清空session
	session.Clear()
	// 保存session数据
	session.Save()
	return nil
}

//GetAdminLoginUid 获取后台登录用户ID
func (svc *Service) GetAdminLoginUid(ctx *gin.Context) int {
	// 初始化session对象
	session := sessions.Default(ctx)
	// 获取用户ID
	userId := gconv.Int(session.Get("adminUserId"))
	// 返回用户ID
	return userId
}

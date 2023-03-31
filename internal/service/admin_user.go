package service

import (
	"context"
	"errors"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/internal/vo"
	"github.com/mwqnice/oh-admin/pkg/constant"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/utils"
)

//UpdatePwd 更新密码
func (svc *Service) UpdatePwd(ctx context.Context, params *dto.UpdatePwdRequest) error {
	// 查询信息
	user, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Model: &model.Model{ID: params.UserId}})
	if err != nil {
		return err
	}
	if user.Model == nil {
		return errors.New("该用户不存在")
	}
	// 比对旧密码
	oldPwd, err := utils.EncodeMD5(params.OldPassword + user.Salt)
	if err != nil {
		return err
	}
	if oldPwd != user.Password {
		return errors.New("旧密码不正确")
	}
	newPwd, err := utils.EncodeMD5(params.NewPassword + user.Salt)
	if err != nil {
		return err
	}
	// 更新密码
	user.Password = newPwd
	return svc.dao.UpdateAdminUser(ctx, user)
}

//CheckAdminPwd 校验密码
func (svc *Service) CheckAdminPwd(ctx context.Context, params *dto.CheckPwdRequest) error {
	// 查询信息
	user, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Model: &model.Model{ID: params.UserId}})
	if err != nil {
		return err
	}
	if user.Model == nil {
		return errors.New("该用户不存在")
	}
	// 比对旧密码
	pwd, err := utils.EncodeMD5(params.Password + user.Salt)
	if err != nil {
		return err
	}
	if pwd != user.Password {
		return errors.New("密码不正确")
	}
	return nil
}

//GetAdminUserInfo 获取用户信息
func (svc *Service) GetAdminUserInfo(ctx context.Context, userId int) (*vo.AdminUserInfoVo, error) {
	// 查询信息
	userInfo, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Model: &model.Model{ID: userId}})
	if err != nil {
		return nil, err
	}
	userInfoVo := &vo.AdminUserInfoVo{}
	userInfoVo.AdminUser = userInfo
	// 性别
	if userInfo.Gender > 0 {
		userInfoVo.GenderName = constant.GENDER_LIST[userInfo.Gender]
	}

	// 角色列表
	roleList, _ := svc.dao.GetAdminUserRoleList(ctx, userInfo.ID)
	if len(roleList) > 0 {
		userInfoVo.RoleList = roleList
	} else {
		userInfoVo.RoleList = make([]model.AdminRole, 0)
	}
	roleIds := make([]int, 0)
	for _, v := range roleList {
		roleIds = append(roleIds, v.ID)
	}
	userInfoVo.RoleIds = roleIds

	return userInfoVo, nil
}

//UpdateAdminUser 更新用户
func (svc *Service) UpdateAdminUser(ctx context.Context, params *dto.UpdateAdminUserRequest) error {
	// 查询信息
	adminUser, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Model: &model.Model{ID: convert.Int(params.Id)}})
	if err != nil {
		return err
	}
	if adminUser.Model == nil {
		return errors.New("该用户不存在")
	}
	//判断用户名是否有重复
	isExist, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Username: params.Username})
	if err != nil {
		return err
	}
	if isExist.Model != nil && isExist.Model.ID != adminUser.ID {
		return errors.New("该用户已经存在")
	}

	adminUser.Realname = params.Realname
	adminUser.Gender = convert.Int(params.Gender)
	adminUser.Mobile = params.Mobile
	adminUser.Email = params.Email
	adminUser.Username = params.Username
	adminUser.Status = convert.Int(params.Status)
	adminUser.Intro = params.Intro
	adminUser.Address = params.Address

	if params.Password != "" {
		pwd, _ := utils.EncodeMD5(params.Password + adminUser.Salt)
		adminUser.Password = pwd
	}
	// 头像处理
	if params.Avatar != "" {
		avatar, err := utils.SaveImage(params.Avatar, "user")
		if err != nil {
			return err
		}
		adminUser.Avatar = avatar
	}

	//更新用户
	err = svc.dao.UpdateAdminUser(ctx, adminUser)
	if err != nil {
		return err
	}

	//删除用户角色
	err = svc.dao.DeleteAdminUserRole(ctx, &model.AdminUserRole{UserId: adminUser.ID})
	if err != nil {
		return err
	}
	if params.RoleIds != "" {
		var userRole []*model.AdminUserRole
		roleIds := convert.ToInt64Array(params.RoleIds, ",")
		for _, val := range roleIds {
			userRole = append(userRole, &model.AdminUserRole{
				UserId: adminUser.ID,
				RoleId: int(val),
			})
		}

		svc.dao.BatchCreateAdminUserRole(ctx, userRole)
	}

	return nil
}

//GetAdminUserList 获取用户列表
func (svc *Service) GetAdminUserList(ctx context.Context, params *dto.GetAdminUserListRequest) ([]*vo.AdminUserInfoVo, int64, error) {
	list, total, err := svc.dao.GetAdminUserList(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	// 数据处理
	var result []*vo.AdminUserInfoVo
	for _, v := range list {
		item := &vo.AdminUserInfoVo{}
		item.AdminUser = v
		// 性别
		if v.Gender > 0 {
			item.GenderName = constant.GENDER_LIST[v.Gender]
		}

		// 角色列表
		roleList, _ := svc.dao.GetAdminUserRoleList(ctx, v.ID)
		if len(roleList) > 0 {
			item.RoleList = roleList
		} else {
			item.RoleList = make([]model.AdminRole, 0)
		}
		// 加入数组
		result = append(result, item)
	}
	return result, total, nil
}

//CreateAdminUser 创建后台用户
func (svc *Service) CreateAdminUser(ctx context.Context, params *dto.CreateAdminUserRequest) error {
	//判断该用户是否存在
	isExist, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Username: params.Username})
	if err != nil {
		return err
	}
	if isExist.Model != nil {
		return errors.New("该用户已经存在")
	}
	adminUser := &model.AdminUser{
		Realname: params.Realname,
		Gender:   convert.Int(params.Gender),
		Avatar:   params.Avatar,
		Mobile:   params.Mobile,
		Email:    params.Email,
		Username: params.Username,
		Status:   convert.Int(params.Status),
		Intro:    params.Intro,
		Address:  params.Address,
	}
	if params.Password != "" {
		salt := utils.GetRandomString(6)
		pwd, _ := utils.EncodeMD5(params.Password + salt)
		adminUser.Password = pwd
		adminUser.Salt = salt
	}
	// 头像处理
	if params.Avatar != "" {
		avatar, err := utils.SaveImage(params.Avatar, "user")
		if err != nil {
			return err
		}
		adminUser.Avatar = avatar
	}

	//创建用户
	user, err := svc.dao.CreateAdminUser(ctx, adminUser)
	if err != nil {
		return err
	}

	if params.RoleIds != "" {
		var userRole []*model.AdminUserRole
		roleIds := convert.ToInt64Array(params.RoleIds, ",")
		for _, val := range roleIds {
			userRole = append(userRole, &model.AdminUserRole{
				UserId: user.ID,
				RoleId: int(val),
			})
		}

		svc.dao.BatchCreateAdminUserRole(ctx, userRole)
	}

	return nil
}

//SetAdminUserStatus 设置后台用户状态
func (svc *Service) SetAdminUserStatus(ctx context.Context, params *dto.SetAdminUserStatusRequest) error {
	adminUser, err := svc.dao.GetAdminUser(ctx, &model.AdminUser{Model: &model.Model{ID: convert.Int(params.Id)}})
	if err != nil {
		return err
	}
	if adminUser.Model == nil {
		return errors.New("该记录不存在")
	}
	adminUser.Status = convert.Int(params.Status)
	return svc.dao.UpdateAdminUser(ctx, adminUser)
}

//DeleteAdminUser 删除后台用户
func (svc *Service) DeleteAdminUser(ctx context.Context, id string) error {
	return svc.dao.DeleteAdminUser(ctx, convert.Int64(id))
}

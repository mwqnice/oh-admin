package service

import (
	"context"
	"errors"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/internal/vo"
	"github.com/mwqnice/oh-admin/pkg/convert"
)

//GetAdminRoleList 获取后台角色列表
func (svc *Service) GetAdminRoleList(ctx context.Context, params *dto.GetAdminRoleListRequest) ([]*model.AdminRole, int64, error) {
	return svc.dao.GetAdminRoleList(ctx, params)
}

//CreateAdminRole 创建后台角色
func (svc *Service) CreateAdminRole(ctx context.Context, params *dto.CreateAdminRoleRequest) error {
	adminRole := &model.AdminRole{}
	convert.SwapTo(params, &adminRole)
	return svc.dao.CreateAdminRole(ctx, adminRole)
}

//GetAdminRoleInfo 获取后台角色详情
func (svc *Service) GetAdminRoleInfo(ctx context.Context, id int64) (*model.AdminRole, error) {
	return svc.dao.GetAdminRoleInfo(ctx, id)
}

//UpdateAdminRole 更新后台角色
func (svc *Service) UpdateAdminRole(ctx context.Context, params *dto.UpdateAdminRoleRequest) error {
	adminRole, err := svc.dao.GetAdminRoleInfo(ctx, convert.Int64(params.Id))
	if err != nil {
		return err
	}
	if adminRole.Model == nil {
		return errors.New("该记录不存在")
	}
	adminRole.Name = params.Name
	adminRole.Code = params.Code
	adminRole.Status = convert.Int(params.Status)
	adminRole.Sort = convert.Int(params.Sort)
	return svc.dao.UpdateAdminRole(ctx, adminRole)
}

//DeleteAdminRole 删除后台角色
func (svc *Service) DeleteAdminRole(ctx context.Context, ids string) error {
	idsArr := convert.ToInt64Array(ids, ",")
	return svc.dao.DeleteAdminRole(ctx, idsArr)
}

//GetRoleMenuList 获取角色菜单列表
func (svc *Service) GetRoleMenuList(ctx context.Context, roleId string) ([]*vo.RoleMenuInfo, error) {
	// 获取全部菜单列表
	adminMenuList, adminMenuTotal, err := svc.dao.GetAdminMenuList(ctx, &dto.GetAdminMenuListRequest{})
	if err != nil {
		return nil, err
	}
	if adminMenuTotal == 0 {
		return nil, errors.New("菜单列表不存在")
	}
	// 获取角色菜单权限列表
	adminRoleMenuList, err := svc.dao.GetAdminRoleMenuList(ctx, &model.AdminRoleMenu{RoleId: convert.Int(roleId)})
	if err != nil {
		return nil, err
	}
	idList := make(map[int]int, 0)
	for _, v := range adminRoleMenuList {
		idList[v.MenuId] = v.MenuId
	}
	// 对象处理
	var list []*vo.RoleMenuInfo
	if len(adminMenuList) > 0 {
		for _, m := range adminMenuList {
			var info = &vo.RoleMenuInfo{
				Id:   m.ID,
				Name: m.Name,
				Open: true,
				Pid:  m.Pid,
			}
			// 节点选中值
			if _, ok := idList[m.Model.ID]; ok {
				info.Checked = true
			}
			list = append(list, info)
		}
	}
	return list, nil
}

//SaveRoleMenuList 保存角色菜单列表
func (svc *Service) SaveRoleMenuList(ctx context.Context, params *dto.SaveAdminRoleMenuRequest) error {
	// 删除现有的角色权限数据
	err := svc.dao.DeleteAdminRoleMenu(ctx, &model.AdminRoleMenu{RoleId: params.RoleId})
	if err != nil {
		return err
	}

	// 遍历创建新角色权限数据
	itemArr := convert.ToInt64Array(params.MenuIds, ",")
	var insertData []*model.AdminRoleMenu
	for _, v := range itemArr {
		insertData = append(insertData, &model.AdminRoleMenu{
			RoleId: params.RoleId,
			MenuId: int(v),
		})
	}
	_, err = svc.dao.BatchCreateAdminRoleMenu(ctx, insertData)
	return err
}

//SetAdminRoleStatus 设置后台角色状态
func (svc *Service) SetAdminRoleStatus(ctx context.Context, params *dto.SetAdminRoleStatusRequest) error {
	adminRole, err := svc.dao.GetAdminRoleInfo(ctx, convert.Int64(params.Id))
	if err != nil {
		return err
	}
	if adminRole.Model == nil {
		return errors.New("该记录不存在")
	}
	adminRole.Status = convert.Int(params.Status)
	return svc.dao.UpdateAdminRole(ctx, adminRole)
}

func (svc *Service) GetAdminRoleAllList(ctx context.Context) ([]*model.AdminRole, error) {
	list, _, err := svc.dao.GetAdminRoleList(ctx, &dto.GetAdminRoleListRequest{Status: "1"})
	return list, err
}

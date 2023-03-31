package service

import (
	"context"
	"errors"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/internal/vo"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"reflect"
)

//GetPermissionMenuList 获取菜单权限列表
func (svc *Service) GetPermissionMenuList(ctx context.Context, userId int) interface{} {
	return svc.dao.GetAdminMenuListByUid(ctx, userId)
}

//GetAdminMenuList 获取后台菜单列表
func (svc *Service) GetAdminMenuList(ctx context.Context, params *dto.GetAdminMenuListRequest) ([]*model.AdminMenu, int64, error) {
	return svc.dao.GetAdminMenuList(ctx, params)
}

//GetAdminMenuTreeList 获取菜单树列表
func (svc *Service) GetAdminMenuTreeList(ctx context.Context) ([]*vo.MenuTreeNode, error) {
	return svc.dao.GetAdminMenuTreeList(ctx)
}

//MakeMenuTreeList 数据源转换
func (svc *Service) MakeMenuTreeList(data []*vo.MenuTreeNode) map[int]string {
	menuList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			menuList[val.ID] = val.Name

			// 二级栏目
			for _, v := range val.Children {
				menuList[v.ID] = "|--" + v.Name

				// 三级栏目
				for _, vt := range v.Children {
					menuList[vt.ID] = "|--|--" + vt.Name
				}
			}
		}
	}
	return menuList
}

//CreateAdminMenu 创建后台菜单
func (svc *Service) CreateAdminMenu(ctx context.Context, params *dto.CreateAdminMenuRequest) error {

	//判断该菜单是否已经存在
	isExistMenu, err := svc.dao.GetAdminMenuInfo(ctx, &model.AdminMenu{Permission: params.Permission})
	if err != nil {
		return err
	}
	if isExistMenu.Model != nil {
		return errors.New("该权限标识已经存在")
	}

	adminMenu := &model.AdminMenu{
		Name:       params.Name,
		Icon:       params.Icon,
		Url:        params.Url,
		Pid:        convert.Int(params.Pid),
		Type:       convert.Int(params.Type),
		IsShow:     convert.Int(params.IsShow),
		Sort:       convert.Int(params.Sort),
		Permission: params.Permission,
		Remark:     params.Remark,
		Status:     convert.Int(params.Status),
	}
	return svc.dao.CreateAdminMenu(ctx, adminMenu)
}

//GetAdminMenuInfo 获取后台菜单详情
func (svc *Service) GetAdminMenuInfo(ctx context.Context, id int64) (*model.AdminMenu, error) {
	return svc.dao.GetAdminMenuInfo(ctx, &model.AdminMenu{Model: &model.Model{ID: int(id)}})
}

//UpdateAdminMenu 更新后台菜单
func (svc *Service) UpdateAdminMenu(ctx context.Context, params *dto.UpdateAdminMenuRequest) error {
	adminMenu, err := svc.dao.GetAdminMenuInfo(ctx, &model.AdminMenu{Model: &model.Model{ID: convert.Int(params.Id)}})
	if err != nil {
		return err
	}
	if adminMenu.Model == nil {
		return errors.New("该记录不存在")
	}
	adminMenu.Name = params.Name
	adminMenu.Icon = params.Icon
	adminMenu.Url = params.Url
	adminMenu.Pid = convert.Int(params.Pid)
	adminMenu.Type = convert.Int(params.Type)
	adminMenu.IsShow = convert.Int(params.IsShow)
	adminMenu.Remark = params.Remark
	adminMenu.Permission = params.Permission
	adminMenu.Status = convert.Int(params.Status)
	adminMenu.Sort = convert.Int(params.Sort)
	return svc.dao.UpdateAdminMenu(ctx, adminMenu)
}

//DeleteAdminMenu 删除后台菜单
func (svc *Service) DeleteAdminMenu(ctx context.Context, id string) error {
	return svc.dao.DeleteAdminMenu(ctx, convert.Int64(id))
}

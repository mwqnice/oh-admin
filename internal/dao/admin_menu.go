package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/internal/vo"
)

//递归生成分类列表
func makeTree(menu []*model.AdminMenu, tn *vo.MenuTreeNode) {
	for _, c := range menu {
		if (tn.AdminMenu == nil && c.Pid == 0) || (tn.AdminMenu != nil && c.Pid == tn.Model.ID) {
			child := &vo.MenuTreeNode{}
			child.AdminMenu = c
			tn.Children = append(tn.Children, child)
			makeTree(menu, child)
		}
	}
}

//GetAdminMenuList 获取菜单列表
func (d *Dao) GetAdminMenuList(ctx context.Context, params *dto.GetAdminMenuListRequest) ([]*model.AdminMenu, int64, error) {
	adminMenu := model.AdminMenu{}
	return adminMenu.List(ctx, d.engine, params)
}

//GetAdminMenuListByUid 通过用户id获取菜单列表
func (d *Dao) GetAdminMenuListByUid(ctx context.Context, userId int) interface{} {
	adminMenu := &model.AdminMenu{}
	var list []*model.AdminMenu
	d.engine.WithContext(ctx).Table(adminMenu.TableName()+" menu").
		Joins("left join oh_admin_role_menu as rm on menu.id = rm.menu_id").
		Joins("left join oh_admin_user_role as ur on rm.role_id = ur.role_id").
		Order("menu.sort asc, menu.id asc").
		Where("ur.user_id=? AND menu.type=0 AND menu.`status`=1 AND menu.is_show=1", userId).
		Select("menu.*").
		Find(&list)
	var menuNode vo.MenuTreeNode
	makeTree(list, &menuNode)

	return menuNode.Children
}

//GetAdminMenuTreeList 获取菜单树列表
func (d *Dao) GetAdminMenuTreeList(ctx context.Context) ([]*vo.MenuTreeNode, error) {
	adminMenu := model.AdminMenu{}
	list, _, err := adminMenu.List(ctx, d.engine, &dto.GetAdminMenuListRequest{Status: model.STATE_OPEN})
	if err != nil {
		return nil, err
	}
	var menuNode vo.MenuTreeNode
	makeTree(list, &menuNode)
	return menuNode.Children, nil
}

//CreateAdminMenu 创建菜单
func (d *Dao) CreateAdminMenu(ctx context.Context, menu *model.AdminMenu) error {
	return menu.Create(ctx, d.engine)
}

//GetAdminMenuInfo 获取菜单信息
func (d *Dao) GetAdminMenuInfo(ctx context.Context, adminMenu *model.AdminMenu) (*model.AdminMenu, error) {
	return adminMenu.Get(ctx, d.engine)
}

//UpdateAdminMenu 更新菜单
func (d *Dao) UpdateAdminMenu(ctx context.Context, menu *model.AdminMenu) error {
	return menu.Update(ctx, d.engine)
}

//DeleteAdminMenu 删除菜单
func (d *Dao) DeleteAdminMenu(ctx context.Context, id int64) error {
	adminMenu := model.AdminMenu{}
	return adminMenu.Delete(ctx, d.engine, int(id))
}

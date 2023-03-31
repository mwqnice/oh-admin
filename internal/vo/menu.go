package vo

import "github.com/mwqnice/oh-admin/internal/model"

//MenuTreeNode 菜单Vo
type MenuTreeNode struct {
	*model.AdminMenu
	Children []*MenuTreeNode `json:"children"` // 子菜单
}

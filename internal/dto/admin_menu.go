/**
 * @describe menuDto
 * @author mwqnice
 * @since 2023/9/9
 * @File : menu
 */
package dto

//GetAdminMenuListRequest 获取菜单列表
type GetAdminMenuListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

//CreateAdminMenuRequest 添加菜单
type CreateAdminMenuRequest struct {
	Name       string `form:"name" binding:"required" json:"name"`             // 菜单标题
	Icon       string `form:"icon" json:"icon"`                                // 图标
	Url        string `form:"url" binding:"required" json:"url"`               // URL地址
	Pid        string `form:"pid" json:"pid"`                                  // 上级ID
	Type       string `form:"type" json:"type"`                                // 类型：1模块 2导航 3菜单 4节点
	Permission string `form:"permission" binding:"required" json:"permission"` // 权限标识
	Status     string `form:"status" binding:"required" json:"status"`         // 状态：1正常 2禁用
	IsShow     string `form:"is_show" json:"is_show"`                          // 是否显示：1显示 2隐藏
	Remark     string `form:"remark" json:"remark"`                            // 菜单备注
	Sort       string `form:"sort" json:"sort"`                                // 显示顺序
}

//UpdateAdminMenuRequest 更新菜单
type UpdateAdminMenuRequest struct {
	Id         string `form:"id" binding:"required" json:"id"`
	Name       string `form:"name" binding:"required" json:"name"`             // 菜单标题
	Icon       string `form:"icon" json:"icon"`                                // 图标
	Url        string `form:"url" binding:"required" json:"url"`               // URL地址
	Pid        string `form:"pid" json:"pid"`                                  // 上级ID
	Type       string `form:"type" json:"type"`                                // 类型：1模块 2导航 3菜单 4节点
	Permission string `form:"permission" binding:"required" json:"permission"` // 权限标识
	Status     string `form:"status" binding:"required" json:"status"`         // 状态：1正常 2禁用
	IsShow     string `form:"is_show" json:"is_show"`                          // 是否显示：1显示 2隐藏
	Remark     string `form:"remark" json:"remark"`                            // 菜单备注
	Sort       string `form:"sort" json:"sort"`                                // 显示顺序
}

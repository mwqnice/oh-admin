/**
 * @describe linkDto
 * @author mwqnice
 * @since 2023/9/9
 * @File : link
 */
package dto

//GetLinkListRequest 获取友链列表
type GetLinkListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

//CreateLinkRequest 添加友链
type CreateLinkRequest struct {
	Name   string `form:"name" binding:"required" json:"name"`     // 菜单标题
	Url    string `form:"url" binding:"required" json:"url"`       // URL地址
	Image  string `form:"image" json:"image"`                      // 图片
	Status string `form:"status" binding:"required" json:"status"` // 状态：1正常 2禁用
	Sort   string `form:"sort" json:"sort"`                        // 显示顺序
}

//UpdateLinkRequest 更新友链
type UpdateLinkRequest struct {
	Id     string `form:"id" binding:"required" json:"id"`
	Name   string `form:"name" binding:"required" json:"name"`     // 菜单标题
	Url    string `form:"url" binding:"required" json:"url"`       // URL地址
	Image  string `form:"image" json:"image"`                      // 图片
	Status string `form:"status" binding:"required" json:"status"` // 状态：1正常 2禁用
	Sort   string `form:"sort" json:"sort"`                        // 显示顺序
}

//SetLinkStatusRequest 设置状态
type SetLinkStatusRequest struct {
	Id     string `form:"id" binding:"required" json:"id"`
	Status string `form:"status"    binding:"required" json:"status"`
}

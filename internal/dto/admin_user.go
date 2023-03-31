package dto

//GetAdminMenuListRequest 获取菜单列表
type GetAdminUserListRequest struct {
	Name     string `form:"name" json:"name"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

//UpdatePwdRequest 更新密码请求
type UpdatePwdRequest struct {
	OldPassword string `form:"old_password" binding:"required" json:"old_password"` // 旧密码
	NewPassword string `form:"new_password" binding:"required" json:"new_password"` // 新密码
	RePassword  string `form:"re_password" binding:"required" json:"re_password"`   // 确认密码
	UserId      int    `json:"user_id"`
}

//CheckPwdRequest 更新密码请求
type CheckPwdRequest struct {
	Password string `form:"password" binding:"required" json:"password"`
	UserId   int    `json:"user_id"`
}

// 添加用户
type CreateAdminUserRequest struct {
	Realname string `form:"realname" binding:"required" json:"realname"`
	Gender   string `form:"gender" binding:"required" json:"gender"`
	Avatar   string `form:"avatar" binding:"required" json:"avatar"`
	Mobile   string `form:"mobile" binding:"required" json:"mobile"`
	Email    string `form:"email" binding:"required" json:"email"`
	Username string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
	Intro    string `form:"intro" json:"intro"`
	Status   string `form:"status" json:"status"`
	RoleIds  string `form:"role_ids" json:"role_ids"` // 用户角色
}

// 更新用户
type UpdateAdminUserRequest struct {
	Id       string `form:"id" binding:"required" json:"id"`
	Realname string `form:"realname" binding:"required" json:"realname"`
	Gender   string `form:"gender" binding:"required" json:"gender"`
	Avatar   string `form:"avatar" binding:"required" json:"avatar"`
	Mobile   string `form:"mobile" binding:"required" json:"mobile"`
	Email    string `form:"email" binding:"required" json:"email"`
	Username string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
	Intro    string `form:"intro" json:"intro"`
	Status   string `form:"status" json:"status"`
	RoleIds  string `form:"role_ids" json:"role_ids"` // 用户角色
}

// 设置状态
type SetAdminUserStatusRequest struct {
	Id     string `form:"id" binding:"required"`
	Status string `form:"status" binding:"required"`
}

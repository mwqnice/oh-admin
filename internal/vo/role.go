package vo

//RoleMenuInfo 角色权限菜单列表
type RoleMenuInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Pid     int    `json:"pId"`
	Checked bool   `json:"checked"`
	Open    bool   `json:"open"`
}

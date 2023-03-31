package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
)

type indexHandler struct {
	svc service.Service
}

//IndexHandler Public管理对象
var IndexHandler = new(indexHandler)

//Index 首页
func (c *indexHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	adminLoginUid := svc.GetAdminLoginUid(ctx)

	// 获取用户信息
	userInfo, _ := svc.GetAdminUserInfo(svc.GetCtx(), adminLoginUid)

	// 获取用户授权菜单
	menuList := svc.GetPermissionMenuList(svc.GetCtx(), adminLoginUid)

	// 渲染模板并绑定数据
	response.BuildTpl(ctx, "index.html").WriteTpl(gin.H{
		"userInfo": userInfo,
		"menuList": menuList,
	})
}

//Welcome 欢迎页
func (c *indexHandler) Welcome(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	// 渲染模板并绑定数据
	response.BuildTpl(ctx, "welcome.html").WriteTpl()
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/errcode"
	"github.com/mwqnice/oh-admin/pkg/utils"
	"net/http"
)

type menuHandler struct {
	svc service.Service
}

//MenuHandler 菜单管理对象
var MenuHandler = new(menuHandler)

//Index 首页
func (c *menuHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "menu_index.html").WriteTpl(gin.H{})
}

func (c *menuHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	// 调用查询列表方法
	list, total, err := svc.GetAdminMenuList(svc.GetCtx(), &dto.GetAdminMenuListRequest{})
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}

//Add 添加
func (c *menuHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.CreateAdminMenuRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.CreateAdminMenu(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(dto.SuccessResponse{
			Msg: "添加成功",
		})
		return
	} else {
		pid := ctx.Query("pid")
		menuTreeList, _ := svc.GetAdminMenuTreeList(svc.GetCtx())

		// 渲染模板
		response.BuildTpl(ctx, "menu_add.html").WriteTpl(gin.H{
			"pid":          convert.Int(pid),
			"menuTreeList": menuTreeList,
		})
	}
}

//Edit 编辑
func (c *menuHandler) Edit(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.UpdateAdminMenuRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}

		err := svc.UpdateAdminMenu(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(dto.SuccessResponse{
			Msg: "修改成功",
		})
		return
	} else {
		id := ctx.Query("id")
		info, _ := svc.GetAdminMenuInfo(svc.GetCtx(), int64(convert.Int(id)))

		menuTreeList, _ := svc.GetAdminMenuTreeList(svc.GetCtx())
		// 渲染模板
		response.BuildTpl(ctx, "menu_edit.html").WriteTpl(gin.H{
			"info":         info,
			"menuTreeList": menuTreeList,
		})
	}
}

//Delete 删除
func (c *menuHandler) Delete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := ctx.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteAdminMenu(svc.GetCtx(), id)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "删除成功",
	})
	return
}

//IndexOld 首页
func (c *menuHandler) IndexOld(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetAdminMenuListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetAdminMenuList(svc.GetCtx(), params)
	if err != nil {
		return
	}
	// 渲染模板并绑定数据
	var pages = utils.NewPage(app.GetPage(ctx), app.GetPageSize(ctx), int(total), "index")
	response.BuildTpl(ctx, "menu_index.html").WriteTpl(gin.H{
		"page":     pages.Show(),
		"menuList": list,
	})
}

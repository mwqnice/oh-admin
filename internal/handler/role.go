package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/errcode"
	"net/http"
)

type roleHandler struct {
	svc service.Service
}

//RoleHandler 角色管理对象
var RoleHandler = new(roleHandler)

//Index 首页
func (c *roleHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "role_index.html").WriteTpl(gin.H{})
}

//List 列表
func (c *roleHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetAdminRoleListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetAdminRoleList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}

//Add 添加
func (c *roleHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	if ctx.Request.Method == http.MethodPost {
		params := &dto.CreateAdminRoleRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		svc := service.New(ctx.Request.Context())
		err := svc.CreateAdminRole(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(dto.SuccessResponse{
			Msg: "添加成功",
		})
		return
	} else {
		// 渲染模板
		response.BuildTpl(ctx, "role_add.html").WriteTpl()
	}
}

//Edit 编辑
func (c *roleHandler) Edit(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.UpdateAdminRoleRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}

		err := svc.UpdateAdminRole(svc.GetCtx(), params)
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
		info, _ := svc.GetAdminRoleInfo(svc.GetCtx(), int64(convert.Int(id)))

		// 渲染模板
		response.BuildTpl(ctx, "role_edit.html").WriteTpl(gin.H{
			"info": info,
		})
	}
}

//Delete 删除
func (c *roleHandler) Delete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	ids := ctx.Param("ids")
	if ids == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteAdminRole(svc.GetCtx(), ids)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "删除成功",
	})
	return
}

//MenuList 角色菜单列表
func (c *roleHandler) MenuList(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	// 角色ID
	roleId := ctx.Param("role_id")
	if roleId == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(ctx.Request.Context())
	list, err := svc.GetRoleMenuList(svc.GetCtx(), roleId)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponseList(list, 0)
	return
}

//SaveRoleMenuList 保存角色菜单列表
func (c *roleHandler) SaveRoleMenuList(ctx *gin.Context) {
	params := &dto.SaveAdminRoleMenuRequest{}
	response := app.NewResponse(ctx)
	// 角色ID
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(ctx.Request.Context())
	err := svc.SaveRoleMenuList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "保存成功",
	})
	return
}

//SetStatus 设置状态
func (c *roleHandler) SetStatus(ctx *gin.Context) {
	params := &dto.SetAdminRoleStatusRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(ctx.Request.Context())
	err := svc.SetAdminRoleStatus(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "设置成功",
	})
	return
}

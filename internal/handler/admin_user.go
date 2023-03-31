package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/constant"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/errcode"
	"net/http"
)

type adminUserHandler struct {
	svc service.Service
}

//AdminUserHandler 用户管理对象
var AdminUserHandler = new(adminUserHandler)

//Index 用户列表页
func (c *adminUserHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	// 渲染模板并绑定数据
	response.BuildTpl(ctx, "user_index.html").WriteTpl()
}
func (c *adminUserHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetAdminUserListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetAdminUserList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}

//UserInfo 用户详情
func (c *adminUserHandler) UserInfo(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.UpdateAdminUserRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		svc.UpdateAdminUser(svc.GetCtx(), params)
		return
	}
	// 获取用户信息
	userInfo, _ := svc.GetAdminUserInfo(svc.GetCtx(), svc.GetAdminLoginUid(ctx))

	// 渲染模板并绑定数据
	response.BuildTpl(ctx, "user_info.html").WriteTpl(gin.H{
		"userInfo": userInfo,
	})
}

//Add 添加
func (c *adminUserHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.CreateAdminUserRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.CreateAdminUser(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(dto.SuccessResponse{
			Msg: "添加成功",
		})
		return
	} else {
		roleList, _ := svc.GetAdminRoleAllList(svc.GetCtx())

		// 渲染模板
		response.BuildTpl(ctx, "user_add.html").WriteTpl(gin.H{
			"roleList":   roleList,
			"genderList": constant.GENDER_LIST,
		})
	}
}

//SetStatus 设置状态
func (c *adminUserHandler) SetStatus(ctx *gin.Context) {
	params := &dto.SetAdminUserStatusRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(ctx.Request.Context())
	err := svc.SetAdminUserStatus(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "设置成功",
	})
	return
}

//Edit 编辑
func (c *adminUserHandler) Edit(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.UpdateAdminUserRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.UpdateAdminUser(svc.GetCtx(), params)
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
		info, _ := svc.GetAdminUserInfo(svc.GetCtx(), convert.Int(id))
		roleList, _ := svc.GetAdminRoleAllList(svc.GetCtx())

		// 渲染模板
		response.BuildTpl(ctx, "user_edit.html").WriteTpl(gin.H{
			"info":       info,
			"roleList":   roleList,
			"genderList": constant.GENDER_LIST,
		})
	}
}

//Delete 删除
func (c *adminUserHandler) Delete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := ctx.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteAdminUser(svc.GetCtx(), id)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "删除成功",
	})
	return
}

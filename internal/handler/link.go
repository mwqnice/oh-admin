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

type linkHandler struct {
	svc service.Service
}

//LinkHandler 友链管理对象
var LinkHandler = new(linkHandler)

//Index 首页
func (c *linkHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "link_index.html").WriteTpl(gin.H{})
}

func (c *linkHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetLinkListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetLinkList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}

//Add 添加
func (c *linkHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.CreateLinkRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.CreateLink(svc.GetCtx(), params)
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
		response.BuildTpl(ctx, "link_add.html").WriteTpl(gin.H{})
	}
}

//Edit 编辑
func (c *linkHandler) Edit(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.UpdateLinkRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}

		err := svc.UpdateLink(svc.GetCtx(), params)
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
		info, _ := svc.GetLinkInfo(svc.GetCtx(), int64(convert.Int(id)))

		// 渲染模板
		response.BuildTpl(ctx, "link_edit.html").WriteTpl(gin.H{
			"info": info,
		})
	}
}

//Delete 删除
func (c *linkHandler) Delete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := ctx.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteLink(svc.GetCtx(), id)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "删除成功",
	})
	return
}

//SetStatus 设置状态
func (c *linkHandler) SetStatus(ctx *gin.Context) {
	params := &dto.SetLinkStatusRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(ctx.Request.Context())
	err := svc.SetLinkStatus(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "设置成功",
	})
	return
}

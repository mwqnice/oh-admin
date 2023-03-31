package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/errcode"
	"github.com/mwqnice/oh-admin/pkg/utils"
	"net/http"
)

type publicHandler struct {
	svc service.Service
}

//PublicHandler Public管理对象
var PublicHandler = new(publicHandler)

//Login 登录
func (c *publicHandler) Login(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	if ctx.Request.Method == http.MethodPost {
		params := &dto.AdminUserLoginRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		svc := service.New(ctx.Request.Context())
		// 校验验证码
		verifyRes := svc.CheckCaptcha(params.IdKey, params.Captcha)
		if !verifyRes {
			response.ToErrorResponse(errcode.CheckCaptchaError.WithDetails(errs.Errors()...))
			return
		}
		//校验登录
		params.Ip = utils.GetClientIp(ctx)
		user, err := svc.AdminUserLogin(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExist.WithDetails(errs.Errors()...))
			return
		}
		// 初始化session对象
		session := sessions.Default(ctx)
		// 设置session数据
		session.Set("adminUserId", user.ID)
		// 保存session数据
		session.Save()
		response.ToResponse(dto.SuccessResponse{
			Msg: "登录成功",
		})
		return
	} else {
		// 渲染模板并绑定数据
		response.BuildTpl(ctx, "login.html").WriteTpl()
	}
}

//Captcha 获取验证码
func (c *publicHandler) Captcha(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	captcha, err := svc.GetCaptcha()
	if err != nil {
		response.ToErrorResponse(errcode.GetCaptchaError.WithDetails(err.Error()))
		return
	}
	response.ToResponse(captcha)
	return
}

//LoginOut 退出登录
func (c *publicHandler) LoginOut(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	err := svc.AdminLoginOut(ctx)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	// 跳转登录页,方式：301(永久移动),308(永久重定向),307(临时重定向)
	ctx.Redirect(http.StatusTemporaryRedirect, "/admin/login")
}

//UpdatePwd 修改密码
func (c *publicHandler) UpdatePwd(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.UpdatePwdRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	if params.NewPassword != params.RePassword {
		response.ToErrorResponse(errcode.Fail.WithDetails("两次密码不一致"))
		return
	}
	params.UserId = ctx.GetInt("adminLoginUid")
	svc := service.New(ctx.Request.Context())
	err := svc.UpdatePwd(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(dto.SuccessResponse{
		Msg: "修改成功",
	})
	return
}

//CheckPwd 校验密码
func (c *publicHandler) CheckPwd(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.CheckPwdRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	params.UserId = ctx.GetInt("adminLoginUid")
	svc := service.New(ctx.Request.Context())
	err := svc.CheckAdminPwd(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(dto.SuccessResponse{})
	return
}

//UploadImage 上传图片
func (c *publicHandler) UploadImage(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	// 调用上传方法
	svc := service.New(ctx.Request.Context())
	result, err := svc.UploadImage(ctx)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(dto.SuccessResponse{
		Code: 0,
		Msg:  "上传成功",
		Data: result,
	})
}

package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/pkg/errcode"
)

const (
	CODE_SUCCESS     = 200
	CODE_DOU_SUCCESS = 0
	MSG_SUCCESS      = "success"
	CODE_ROUTE_ERROR = 404
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

type ResponseCommonStruct struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Pager Pager       `json:"pager"`
	Count int         `json:"count"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, &ResponseCommonStruct{
		Data:  list,
		Count: totalRows,
		Pager: Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(http.StatusOK, response)
}

// 通用tpl响应
type TplResp struct {
	c   *gin.Context
	tpl string
}

//BuildTpl 返回一个tpl响应
func (r *Response) BuildTpl(c *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: tpl,
	}
	return &t
}

//ErrorTpl 返回一个错误的tpl响应
func (r *Response) ErrorTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: "error/error.html",
	}
	return &t
}

//WriteTpl 输出页面模板附加自定义函数
func (resp *TplResp) WriteTpl(params ...gin.H) {
	if params == nil || len(params) == 0 {
		resp.c.HTML(http.StatusOK, resp.tpl, gin.H{})
	} else {
		resp.c.HTML(http.StatusOK, resp.tpl, params[0])
	}
}

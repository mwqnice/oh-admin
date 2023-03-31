package router

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/handler"
	"github.com/mwqnice/oh-admin/internal/middleware"
	"github.com/mwqnice/oh-admin/internal/widget"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/limiter"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"html/template"
	"path/filepath"
	"strings"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("MsW32dQN2342434I5C43E6"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("ohadmin", store))

	//加载模板
	r.HTMLRender = loadTemplates("views")
	// 设置静态资源路由
	r.Static("/static", "./static")

	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	//r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, "OK")
		return
	})
	/* 文件上传 */
	upload := r.Group("upload")
	{
		// 上传图片
		upload.POST("/uploadImage", handler.PublicHandler.UploadImage)
	}

	r.Use(middleware.Tracer())
	r.Use(middleware.AdminAuth()) //验证登录
	index := r.Group("/")
	{
		index.GET("/", handler.IndexHandler.Index)                 //首页
		index.GET("/index", handler.IndexHandler.Index)            //首页
		index.GET("/welcome", handler.IndexHandler.Welcome)        //欢迎页
		index.Any("/login", handler.PublicHandler.Login)           //登录
		index.GET("/captcha", handler.PublicHandler.Captcha)       //获取验证码
		index.GET("/logout", handler.PublicHandler.LoginOut)       //退出
		index.POST("/update_pwd", handler.PublicHandler.UpdatePwd) //修改密码
		index.POST("/check_pwd", handler.PublicHandler.CheckPwd)   //校验密码
		index.Any("/user_info", handler.AdminUserHandler.UserInfo) //获取用户信息

	}
	/* 管理员管理 */
	user := r.Group("user")
	{
		user.GET("/index", handler.AdminUserHandler.Index)          //用户详情
		user.GET("/info", handler.AdminUserHandler.UserInfo)        //用户详情
		user.POST("/list", handler.AdminUserHandler.List)           //菜单列表
		user.Any("/add", handler.AdminUserHandler.Add)              //添加用户
		user.POST("/setStatus", handler.AdminUserHandler.SetStatus) //设置状态
		user.Any("/edit", handler.AdminUserHandler.Edit)            //修改用户
		user.POST("/delete/:id", handler.AdminUserHandler.Delete)   //删除
	}
	/* 菜单管理 */
	menu := r.Group("menu")
	{
		menu.GET("/index", handler.MenuHandler.Index)        //菜单首页
		menu.POST("/list", handler.MenuHandler.List)         //菜单列表
		menu.Any("/add", handler.MenuHandler.Add)            //添加菜单
		menu.Any("/edit", handler.MenuHandler.Edit)          //修改菜单
		menu.POST("/delete/:id", handler.MenuHandler.Delete) //删除
	}
	/* 角色管理 */
	role := r.Group("role")
	{
		role.GET("/index", handler.RoleHandler.Index)                      //角色首页
		role.POST("/list", handler.RoleHandler.List)                       //角色列表
		role.Any("/add", handler.RoleHandler.Add)                          //添加角色
		role.Any("/edit", handler.RoleHandler.Edit)                        //修改角色
		role.POST("/delete/:ids", handler.RoleHandler.Delete)              //删除
		role.POST("/setStatus", handler.RoleHandler.SetStatus)             //设置状态
		role.GET("/menu_list/:role_id", handler.RoleHandler.MenuList)      //角色菜单列表
		role.POST("/menu_list/save", handler.RoleHandler.SaveRoleMenuList) //角色菜单列表保存
	}

	/* 友链管理 */
	link := r.Group("link")
	{
		link.GET("/index", handler.LinkHandler.Index)          //友链首页
		link.POST("/list", handler.LinkHandler.List)           //友链列表
		link.Any("/add", handler.LinkHandler.Add)              //添加友链
		link.Any("/edit", handler.LinkHandler.Edit)            //修改友链
		link.POST("/delete/:id", handler.LinkHandler.Delete)   //删除友链
		link.POST("/setStatus", handler.LinkHandler.SetStatus) //设置状态
	}
	return r
}

func HandleNotFound(c *gin.Context) {
	c.JSON(app.CODE_SUCCESS, &app.ResponseCommonStruct{Code: app.CODE_ROUTE_ERROR, Msg: fmt.Sprintf("路由%s不存在或不支持%s请求", c.Request.URL.String(), c.Request.Method)})
	return
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// 非模板嵌套
	adminHtmls, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, html := range adminHtmls {
		r.AddFromGlob(filepath.Base(html), html)
	}

	// 布局模板
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 嵌套的内容模板
	includes, err := filepath.Glob(templatesDir + "/includes/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// template自定义函数
	funcMap := template.FuncMap{
		"StringToLower": func(str string) string {
			return strings.ToLower(str)
		},
		"date2": func() string {
			return time.Now().Format("2006-01-02 15:04:05.00000")
		},
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
		"query":    widget.Query,
		"add":      widget.Add,
		"edit":     widget.Edit,
		"delete":   widget.Delete,
		"expand":   widget.Expand,
		"collapse": widget.Collapse,
		"addz":     widget.Addz,
		"in":       widget.In,
	}

	// 将主模板，include页面，layout子模板组合成一个完整的html页面
	for _, include := range includes {
		// 文件名称
		baseName := filepath.Base(include)
		files := []string{}
		if strings.Contains(baseName, "edit") || strings.Contains(baseName, "add") {
			files = append(files, templatesDir+"/layouts/form.html", include)
		} else {
			files = append(files, templatesDir+"/layouts/layout.html", include)
		}
		files = append(files, layouts...)
		r.AddFromFilesFuncs(baseName, funcMap, files...)
	}
	return r
}

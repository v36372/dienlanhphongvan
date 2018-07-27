package handler

import (
	"dienlanhphongvan-cdn/client"
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/config"
	"dienlanhphongvan/middleware"
	"dienlanhphongvan/utilities/ulog"
	"path"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func InitEngine(conf *config.Config) *gin.Engine {
	if conf.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware(conf.App.Whitelist))
	r.Use(gin.LoggerWithWriter(ulog.Logger().Request))
	r.Static("static", "./public/static")

	templateConfig := gintemplate.TemplateConfig{
		Root:      "public",
		Extension: ".html",
		Master:    "layouts/master",
		Partials: []string{
			"partials/account",
			"partials/banner",
			"partials/bottom-header",
			"partials/breadcrumbs",
			"partials/checkout",
			"partials/contact",
			"partials/information",
			"partials/map",
			"partials/nav",
			"partials/products",
			"partials/register",
			"partials/single",
			"partials/start-logo",
			"partials/top-header",
		},
		DisableCache: conf.App.Debug,
	}

	r.HTMLRender = gintemplate.New(templateConfig)

	if conf.App.Debug {
		r.Use(gin.Logger())
	}

	var (
		originalImageDir = path.Join(conf.Resource.RootDir, "images")
		cachedImageDir   = path.Join(conf.Resource.RootDir, "cached", "images")
		uploadImageDir   = path.Join(conf.Resource.RootDir, "tmp")
		imgx             = client.NewClient(conf.Imgx.Address, nil)
	)

	// Setup auth middleware
	secCookie := middleware.NewSetCookie(conf.CookieToken.BlockKey, conf.CookieToken.HashKey)
	authMiddleware := middleware.NewAuthMiddleware(secCookie, middleware.Auth.GetLoggedInUser)
	middleware.InitAuth(authMiddleware.GetCurrentUser)

	r.Use(authMiddleware.Interception())

	// INIT ENTITY
	productEntity := entity.NewProduct()
	categoryEntity := entity.NewCategory()
	imageEntity := entity.NewImage(imgx, uploadImageDir, originalImageDir, cachedImageDir, conf.App.Debug)
	userEntity := entity.NewUser()

	// INIT HANDLER
	indexHandler := indexHandler{
		Category: categoryEntity,
		Product:  productEntity,
	}
	productHandler := productHandler{
		productEntity: productEntity,
		imageEntity:   imageEntity,
	}
	dashboardHandler := dashboardHandler{
		product:  productEntity,
		category: categoryEntity,
		image:    imageEntity,
	}
	categoryHandler := categoryHandler{
		category: categoryEntity,
	}
	userHandler := userHandler{
		user:      userEntity,
		secCookie: secCookie,
	}
	imageHandler := imageHandler{
		imageEntity: imageEntity,
	}
	contactHandler := contactHandler{}

	// INIT ROUTE
	groupIndex := r.Group("")
	{
		GET(groupIndex, "", indexHandler.Index)
	}

	groupContact := r.Group("/contact")
	{
		GET(groupContact, "", contactHandler.ContactPage)
	}

	groupProduct := r.Group("/products")
	{
		GET(groupProduct, "/:slug", productHandler.GetDetail)
		GET(groupProduct, "", productHandler.GetList)
	}

	groupImage := r.Group("/images")
	{
		GET(groupImage, "/original/:name", imageHandler.GetOriginal)
		GET(groupImage, "/cached/:name", imageHandler.GetCached)
	}

	groupCategory := r.Group("/categories")
	{
		GET(groupCategory, "/:slug", productHandler.GetByCategory)
	}

	userGroup := r.Group("/user")
	{
		GET(userGroup, "/login", userHandler.LoginPage)
		POST(userGroup, "/login", userHandler.Login)
	}

	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use(authMiddleware.Interception())
	dashboardGroup.Use(authMiddleware.RequireLogin())
	{
		GET(dashboardGroup, "/create-product", dashboardHandler.CreateProduct)
		GET(dashboardGroup, "/update-product/:slug", dashboardHandler.UpdateProduct)
		GET(dashboardGroup, "/create-category", dashboardHandler.CreateCategory)
		GET(dashboardGroup, "/product-list", dashboardHandler.ListProduct)
		POST(dashboardGroup, "/products", productHandler.Create)
		POST(dashboardGroup, "/products/:slug", productHandler.Update)
		POST(dashboardGroup, "/delete-product/:slug", productHandler.Delete)
		POST(dashboardGroup, "/categories", categoryHandler.Create)
	}

	// Image processing
	groupDashboardImage := dashboardGroup.Group("/images")
	{
		POST(groupDashboardImage, "/upload", imageHandler.Upload)
		POST(groupDashboardImage, "/move/:name", imageHandler.Move)
	}

	return r
}

func GET(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "GET", relativePath, f)
}

func POST(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "POST", relativePath, f)
}

func PUT(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "PUT", relativePath, f)
}

func DELETE(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "DELETE", relativePath, f)
}

func route(group *gin.RouterGroup, method string, relativePath string, f func(*gin.Context)) {
	hanld := middleware.ErrorHandler(group.BasePath() + relativePath)
	switch method {
	case "POST":
		group.POST(relativePath, hanld, f)
	case "GET":
		group.GET(relativePath, hanld, f)
	case "PUT":
		group.PUT(relativePath, hanld, f)
	case "DELETE":
		group.DELETE(relativePath, hanld, f)
	}
}

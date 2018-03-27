package handler

import (
	"dienlanhphongvan-cdn/client"
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/config"
	"dienlanhphongvan/middleware"
	"dienlanhphongvan/utilities/ulog"
	"path"

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
	r.LoadHTMLGlob("public/*.html")
	r.Static("static", "./public/static")

	if conf.App.Debug {
		r.Use(gin.Logger())
	}

	var (
		originalImageDir = path.Join(conf.Resource.RootDir, "images")
		cachedImageDir   = path.Join(conf.Resource.RootDir, "cached", "images")
		uploadImageDir   = path.Join(conf.Resource.RootDir, "tmp")
		imgx             = client.NewClient(conf.Imgx.Address, nil)
	)

	indexHandler := indexHandler{
		Category: entity.NewCategory(),
	}
	groupIndex := r.Group("")
	{
		GET(groupIndex, "", indexHandler.Index)
	}

	productEntity := entity.NewProduct()
	imageEntity := entity.NewImage(imgx, uploadImageDir, originalImageDir, cachedImageDir, conf.App.Debug)

	// Product
	productHandler := productHandler{
		productEntity: productEntity,
		imageEntity:   imageEntity,
	}
	groupProduct := r.Group("/products")
	{
		GET(groupProduct, "/:slug", productHandler.GetDetail)
		DEL(groupProduct, "/:slug", productHandler.DeleteProduct)
		GET(groupProduct, "", productHandler.GetList)
		POST(groupProduct, "", productHandler.Create)
	}

	// Dashboard
	dashboardHandler := dashboardHandler{
		product:  productEntity,
		category: entity.NewCategory(),
		image:    imageEntity,
	}

	authMiddleware := middleware.NewAuthMiddleware(secCookie, middleware.Auth.GetLoggedInUser)

	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use(authMiddleware.Interception())
	dashboardGroup.Use(authMiddleware.RequireLogin())
	{
		GET(dashboardGroup, "/create-product", dashboardHandler.CreateProduct)
		GET(dashboardGroup, "/create-category", dashboardHandler.CreateCategory)
		GET(dashboardGroup, "/list-product", dashboardHandler.ListProduct)
	}

	// Image processing
	imageHandler := imageHandler{
		imageEntity: imageEntity,
	}
	groupImage := r.Group("/images")
	{
		POST(groupImage, "/upload", imageHandler.Upload)
		POST(groupImage, "/move/:name", imageHandler.Move)
		GET(groupImage, "/original/:name", imageHandler.GetOriginal)
		GET(groupImage, "/cached/:name", imageHandler.GetCached)
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

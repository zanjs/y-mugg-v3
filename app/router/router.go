package router

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zanjs/y-mugg-v3/app/controllers"
	mid "github.com/zanjs/y-mugg-v3/app/middleware"
	"github.com/zanjs/y-mugg-v3/app/monitor"
	"github.com/zanjs/y-mugg-v3/config"
)

var (
	appConfig = config.Config.App
	jwtConfig = config.Config.JWT
)

// InitRoute 初始化路由
func InitRoute() {

	e := echo.New()

	e.HTTPErrorHandler = monitor.CustomHTTPErrorHandler
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(mid.ServerHeader)

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	e.GET("/", controllers.GetHome)

	e.POST("/user/add", controllers.UserController{}.Create)

	e.GET("/records/jobs", controllers.JobController{}.SyncQnventoryV1)
	e.GET("/jobs/sync_qnventory", controllers.JobController{}.SyncQnventory)

	v0 := e.Group("/v0")

	v0.GET("/", controllers.CreateTable)

	v1 := e.Group("/v1")
	v1.POST("/login", controllers.PostLogin)

	v1.Use(middleware.JWT([]byte(jwtConfig.Secret)))

	// Users
	v1.GET("/users", controllers.UserController{}.GetAll)
	v1.POST("/users", controllers.UserController{}.Create)
	v1.GET("/users/:id", controllers.UserController{}.Get)
	v1.PUT("/users/:id", controllers.UserController{}.Update)
	v1.DELETE("/users/:id", controllers.UserController{}.Delete)

	// Articles
	v1.GET("/articles", controllers.ArticlesController{}.GetAll)
	v1.POST("/articles", controllers.ArticlesController{}.Create)
	v1.GET("/articles/:id", controllers.ArticlesController{}.Get)
	v1.PUT("/articles/:id", controllers.ArticlesController{}.Update)
	v1.DELETE("/articles/:id", controllers.ArticlesController{}.Delete)

	// Products
	v1.GET("/products", controllers.ProductController{}.GetAll)
	v1.POST("/products", controllers.ProductController{}.Create)
	v1.GET("/products/:id", controllers.ProductController{}.Get)
	v1.PUT("/products/:id", controllers.ProductController{}.Update)
	v1.DELETE("/products/:id", controllers.ProductController{}.Delete)

	// Wareroom
	v1.GET("/warerooms", controllers.WareroomController{}.GetAll)
	v1.POST("/warerooms", controllers.WareroomController{}.Create)
	v1.GET("/warerooms/:id", controllers.WareroomController{}.Get)
	v1.PUT("/warerooms/:id", controllers.WareroomController{}.Update)
	v1.DELETE("/warerooms/:id", controllers.WareroomController{}.Delete)

	// 库存
	v1.GET("/inventorys", controllers.InventoryController{}.GetAll)
	v1.POST("/inventorys", controllers.WareroomController{}.Create)
	v1.GET("/inventorys/:id", controllers.WareroomController{}.Get)
	v1.PUT("/inventorys/:id", controllers.WareroomController{}.Update)
	v1.DELETE("/inventorys/:id", controllers.WareroomController{}.Delete)

	// 库存
	v1.GET("/transports", controllers.TransportController{}.GetAll)

	// 销量记录
	v1.GET("/sales", controllers.SaleController{}.GetAll)
	// v1.POST("/sales", controllers.WareroomController{}.Create)
	// v1.GET("/sales/:id", controllers.WareroomController{}.Get)
	v1.PUT("/sales/:id", controllers.SaleController{}.Update)
	// v1.DELETE("/sales/:id", controllers.WareroomController{}.Delete)

	// 库存销量统计
	v1.GET("/statistics", controllers.SattisticsController{}.WhereTime)
	v1.GET("/statistics_fe", controllers.SattisticsController{}.WhereTimeFEData)
	// qm 库存销量更新
	v1.GET("/records", controllers.AllRecordsPage)
	v1.GET("/records/all", controllers.AllRecords)
	v1.GET("/records/q", controllers.GetRecordWhereTime)
	v1.GET("/records/q2", controllers.AllProductWareroomRecordsTime)
	v1.GET("/records/excel", controllers.AllProductWareroomRecords)
	v1.PUT("/records/:id", controllers.UpdateRecord)
	v1.DELETE("/records/:id", controllers.DeleteRecord)

	// Server
	if err := e.Start(fmt.Sprintf("%s:%s", appConfig.HttpAddr, appConfig.HttpPort)); err != nil {
		e.Logger.Fatal(err.Error())
	}

}

package server

import (
	"core_business/internals/core/ports"
	"core_business/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"strconv"
)

// Injection function allow dependency injection into provided methods
type ginRouter struct {
	router *gin.Engine
}

// NewGinRouter creates an instance of the gin router
func NewGinRouter(r *gin.Engine) ports.IRouter {
	return &ginRouter{
		r,
	}
}

func (g *ginRouter) GET(uri string, fn func(ctx *gin.Context)) {
	g.router.GET(uri, fn)
}

func (g *ginRouter) POST(uri string, fn func(ctx *gin.Context)) {
	g.router.POST(uri, fn)
}

func (g *ginRouter) PUT(uri string, fn func(ctx *gin.Context)) {
	g.router.PUT(uri, fn)
}

func (g *ginRouter) DELETE(uri string, fn func(ctx *gin.Context)) {
	g.router.DELETE(uri, fn)
}

func (g *ginRouter) GROUP(path string) *gin.RouterGroup {
	return g.router.Group(path)
}

func (g *ginRouter) SERVE() error {
	var port int
	if config.Instance.Port != nil {

		p, err := strconv.Atoi(*config.Instance.Port)
		if err != nil {
			log.Fatal("Error parsing port, must be numeric")
		}
		port = p
	}

	g.router.Use(gin.Logger())
	g.router.Use(gin.Recovery())
	if err := g.router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return err
	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := g.router.Run(fmt.Sprintf("0.0.0.0:%d", port)); err != nil {
		fmt.Println("Could not run infrastructure -> ", err)
		return err
	}

	return nil
}

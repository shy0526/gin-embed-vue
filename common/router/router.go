/**
 * @Create on : 2023/4/17
 * @Author: sunnyh
 * @Des:
 */

package router

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/shy0526/gin-embed-vue/common/logger"
	"github.com/shy0526/gin-embed-vue/common/settings"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	_apps := router.Group("/apps")
	_appsV1alpha1 := _apps.Group("/v1alpha1")
	{
		_appsV1alpha1.GET("/config", func(ctx *gin.Context) {
			// get app config
			ctx.JSON(http.StatusOK, gin.H{
				"status": 200,
				"config": settings.Conf,
			})
		})

	}

	router.Use(static.Serve(settings.Conf.StaticUrlPrefix, static.LocalFile(settings.Conf.StaticRoot, false)))
	//router.NoRoute(func(c *gin.Context) {
	//	c.File(staticRoot + "/index.html")
	//})

	return router
}

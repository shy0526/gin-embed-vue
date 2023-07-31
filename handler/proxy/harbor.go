/**
 * @Create on: 2023/7/31
 * @Author: sunnyh
 * @File:  harbor.go
 * @Desc:
 */

package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/shy0526/gin-embed-vue/common/settings"
	"net/http/httputil"
	"net/url"
)

func HandlerHarborServer(c *gin.Context) {
	// harbor reverse proxy
	harborConfig := settings.Conf.HarborConfig
	rp := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: harborConfig.Scheme,
		Host:   harborConfig.Host,
	})

	c.Request.URL.Path = c.Param("proxyPath")
	c.Request.SetBasicAuth(harborConfig.User, harborConfig.Password)
	rp.ServeHTTP(c.Writer, c.Request)
}

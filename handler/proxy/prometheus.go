/**
 * @Create on: 2023/7/31
 * @Author: sunnyh
 * @File:  prometheus
 * @Desc:
 */

package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/shy0526/gin-embed-vue/common/settings"
	"net/http/httputil"
	"net/url"
)

func HandlerPrometheusServer(c *gin.Context) {
	// prometheus reverse proxy
	prometheusConfig := settings.Conf.PrometheusConfig
	rp := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: prometheusConfig.Scheme,
		Host:   prometheusConfig.Host,
	})

	c.Request.URL.Path = c.Param("proxyPath")
	rp.ServeHTTP(c.Writer, c.Request)
}

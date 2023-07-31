/**
 * @Create on: 2023/7/31
 * @Author: sunnyh
 * @File:  elasticsearch
 * @Desc:
 */

package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/shy0526/gin-embed-vue/common/settings"
	"net/http/httputil"
	"net/url"
)

func HandlerElasticSearchServer(c *gin.Context) {
	// es reverse proxy
	esConfig := settings.Conf.ElasticSearchConfig
	rp := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: esConfig.Scheme,
		Host:   esConfig.Host,
	})

	c.Request.URL.Path = c.Param("proxyPath")
	c.Request.SetBasicAuth(esConfig.User, esConfig.Password)
	rp.ServeHTTP(c.Writer, c.Request)
}

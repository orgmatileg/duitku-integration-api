package httpwrap

import (
	"duitku-integration-api/internal/handler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// IHTTPWrap interface
type IHTTPWrap interface {
	Start() error
	InitRoute()
}

type httpWrap struct {
	httpServer *gin.Engine
}

// New httpWrap
func New() IHTTPWrap {
	return &httpWrap{
		httpServer: gin.Default(),
	}
}

func (h *httpWrap) Start() error {

	fmt.Println(viper.GetString("env"), "wew")
	if viper.GetString("env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	h.httpServer.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to DuitKu Integration API")
	})

	return h.httpServer.Run()
}

func (h *httpWrap) InitRoute() {
	h.httpServer.POST("/duitku/callback", handler.DuitKuCallBack)
	h.httpServer.POST("/duitku/redirect", handler.DuitKuRedirect)
	h.httpServer.POST("/duitku/request-transaction", handler.DuitKuRequestTransaction)
	h.httpServer.POST("/duitku/check-transaction", handler.DuitKuCheckTransaction)

}

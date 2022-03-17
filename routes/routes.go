package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//-----路由部分start
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	//-----路由部分end

	return r
}

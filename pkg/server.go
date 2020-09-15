package pkg

import (
	"github.com/davidleitw/chronos/pkg/oauth2"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.Default()

	oauth := router.Group("/oauth2")
	{
		oauth.GET("/google", oauth2.GoogleOauthLogin)
		oauth.GET("/facebook", oauth2.FacebookOauthLogin)
	}

	callback := router.Group("/callback")
	{
		callback.GET("/google", oauth2.GoogleCallBack)
		callback.GET("/facebook", oauth2.FacebookCallBack)
	}

	return router
}

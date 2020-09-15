package oauth2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var facebookConfig *oauth2.Config = &oauth2.Config{
	ClientID:     "ID",
	ClientSecret: "Secret",
	RedirectURL:  "",
	Scopes: []string{
		"email",
		"public_profile",
	},
	Endpoint: facebook.Endpoint,
}

func FacebookOauthLogin(ctx *gin.Context) {
	state := xid.New().String()
	redirectURL := googleConfig.AuthCodeURL(state)

	ctx.Redirect(http.StatusSeeOther, redirectURL)
}

func FacebookCallBack(ctx *gin.Context) {

}

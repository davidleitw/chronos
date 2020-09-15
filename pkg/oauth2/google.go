package oauth2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleConfig *oauth2.Config = &oauth2.Config{
	ClientID:     "ID",
	ClientSecret: "Secret",
	RedirectURL:  "",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func GoogleOauthLogin(ctx *gin.Context) {
	state := xid.New().String()
	redirectURL := googleConfig.AuthCodeURL(state)

	ctx.Redirect(http.StatusSeeOther, redirectURL)
}

func GoogleCallBack(ctx *gin.Context) {

}

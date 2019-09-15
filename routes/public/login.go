package publicroutes

import (
	"crypto/rand"
	"encoding/base64"
	"enterprise/app"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func login(c *gin.Context) {
	session := sessions.Default(c)
	domain := app.Config["Auth"]["Domain"]
	aud := app.Config["Auth"]["Audience"]

	conf := &oauth2.Config{
		ClientID:     app.Config["Auth"]["ClientID"],
		ClientSecret: app.Config["Auth"]["Secret"],
		RedirectURL:  app.Config["Auth"]["RedirectURL"],
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	if aud == "" {
		aud = "https://" + domain + "/userinfo"
	}

	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)

	audience := oauth2.SetAuthURLParam("audience", aud)
	url := conf.AuthCodeURL(state, audience)
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, url)
}

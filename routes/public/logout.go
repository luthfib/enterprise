package publicroutes

import (
	"enterprise/app"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {
	domain := app.Config["Auth"]["Domain"]
	session := sessions.Default(c)

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)

	if err != nil {
		panic("boom")
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", "http://localhost:8080")
	parameters.Add("client_id", app.Config["Auth"]["ClientID"])
	Url.RawQuery = parameters.Encode()
	session.Clear()
	session.Save()

	http.Redirect(c.Writer, c.Request, Url.String(), http.StatusTemporaryRedirect)
}

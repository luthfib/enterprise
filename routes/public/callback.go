package publicroutes

import (
	"context"
	_ "crypto/sha512"
	"fmt"
	"net/http"

	"encoding/json"

	"enterprise/app"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func callback(c *gin.Context) {
	session := sessions.Default(c)
	domain := app.Config["Auth"]["Domain"]
	// user := session.Get("user")

	conf := &oauth2.Config{
		ClientID:     app.Config["Auth"]["ClientID"],
		ClientSecret: app.Config["Auth"]["ClientSecret"],
		RedirectURL:  app.Config["Auth"]["RedirectURL"],
		Scopes:       []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	// session_state := session.Get("state")
	// state := c.Query("state")
	// fmt.Println(state)

	// session, err := app.Store.Get(r, "state")

	/*if state != session_state {
		http.Error(c.Writer, "Invalid state parameter", http.StatusInternalServerError)
		return
	}*/

	code := c.Query("code")

	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	client := conf.Client(context.TODO(), token)
	resp, err := client.Get("https://" + domain + "/userinfo")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var profile map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Set("id_token", token.Extra("id_token"))
	session.Set("access_token", token.AccessToken)
	session.Set("user", "Luthfi")
	session.Set("profile", profile)
	sess_err := session.Save()
	if sess_err != nil {
		fmt.Println("error in callback", sess_err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
	}

	c.Redirect(http.StatusSeeOther, "/")
}

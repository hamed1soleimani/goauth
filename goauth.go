package main

import (
	"fmt"

	"github.com/hamed1soleimani/goauth/social"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/sessions"
)


func main() {
	m := martini.Classic()
	m.Use(sessions.Sessions("my_session", sessions.NewCookieStore([]byte("secret123"))))

	m.Use(social.GoogleOauthConfig("oauth.ini"))

	m.Get("/", oauth2.LoginRequired, func(tokens oauth2.Tokens) string {
		if tokens.Expired() {
			return "not logged in, or the access token is expired"
		}
		fmt.Println(tokens.Access())
		return social.GoogleProfileJson(tokens.Access())
	})

	m.Get("/restrict", oauth2.LoginRequired, func(tokens oauth2.Tokens) string {
		return tokens.Access()
	})

	m.Run()
}

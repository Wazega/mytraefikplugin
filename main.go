package mytraefikplugin

import (
	"context"
	"net/http"
)

type Config struct {
	RedirectURL string `json:"redirectURL,omitempty"`
	CookieName  string `json:"cookieName,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		RedirectURL: "http://monApp.localhost/monApp-api/Login",
		CookieName:  "authtoken",
	}
}

type MyTraefikPlugin struct {
	next        http.Handler
	redirectURL string
	cookieName  string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &MyTraefikPlugin{
		next:        next,
		redirectURL: config.RedirectURL,
		cookieName:  config.CookieName,
	}, nil
}

func (a *MyTraefikPlugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(a.cookieName)

	if err != nil {
		// Cookie absent → redirection
		http.Redirect(w, r, a.redirectURL, http.StatusFound) // 302
		return
	}

	// Cookie présent → on laisse passer
	a.next.ServeHTTP(w, r)
}

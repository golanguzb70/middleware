package basicauth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// method for checking authorization
func (cfg *Config) Middleware(ctx *gin.Context) {
	var (
		authRequired = false
		url          = ctx.Request.URL.Path
		method       = ctx.Request.Method
		authHeader   = ctx.GetHeader("Authorization")
	)
	fmt.Println("URL: ", url)
	fmt.Println("Method: ", method)

	if strings.Contains(strings.Join(cfg.RestrictedMethods, ","), method) {
		authRequired = true
	}

	if !authRequired && len(cfg.RestrictedUrls) > 0 {
		for _, e := range cfg.RestrictedUrls {
			if strings.Contains(e, "*") && strings.Contains(url, strings.TrimSuffix(e, "/*")) {
				authRequired = true
				break
			} else if strings.Contains(e, "{") && e == string(url[:strings.LastIndex(url, "/")]) {
				authRequired = true
				break
			} else if e == url {
				authRequired = true
			}
		}
	}

	if authRequired {
		for _, u := range cfg.Users {
			if authHeader == "" {
				ctx.Header("WWW-Authenticate", "Basic realm=Authorization Required")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			credentials := strings.SplitN(authHeader, " ", 2)
			if len(credentials) != 2 {
				ctx.Header("WWW-Authenticate", "Basic realm=Authorization Required")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			decodedCredentials, err := base64.StdEncoding.DecodeString(credentials[1])
			if err != nil {
				ctx.Header("WWW-Authenticate", "Basic realm=Authorization Required")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			credentialsPair := strings.SplitN(string(decodedCredentials), ":", 2)
			if len(credentialsPair) != 2 {
				ctx.Header("WWW-Authenticate", "Basic realm=Authorization Required")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			if credentialsPair[0] != u.UserName || credentialsPair[1] != u.Password {
				ctx.Header("WWW-Authenticate", "Basic realm=Authorization Required")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
	}
	ctx.Next()
}

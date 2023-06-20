package basicauth

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)
// method for checking authorization
func Middleware(cfg Config) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				authRequired = cfg.RequireAuthForAll
				url          = r.URL.Path
				method       = r.Method
				authHeader   = r.Header.Get("Authorization")
			)

			if contains(cfg.RestrictedMethods, method) {
				authRequired = true
			}

			if !authRequired && len(cfg.RestrictedUrls) > 0 {
				for _, e := range cfg.RestrictedUrls {
					if strings.Contains(e, "*") && strings.Contains(url, strings.TrimSuffix(e, "/*")) {
						authRequired = true
						break
					} else if strings.Contains(e, "{") && e[:strings.LastIndex(e, "/")] == url[:strings.LastIndex(url, "/")] {
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
						w.Header().Set("WWW-Authenticate", "Basic realm=Authorization Required")
						cfg.UnauthorizedHandler(w, r)
						return
					}

					credentials := strings.SplitN(authHeader, " ", 2)
					if len(credentials) != 2 {
						w.Header().Set("WWW-Authenticate", "Basic realm=Authorization Required")
						cfg.UnauthorizedHandler(w, r)
						return
					}

					decodedCredentials, err := base64.StdEncoding.DecodeString(credentials[1])
					if err != nil {
						w.Header().Set("WWW-Authenticate", "Basic realm=Authorization Required")
						cfg.UnauthorizedHandler(w, r)
						return
					}

					credentialsPair := strings.SplitN(string(decodedCredentials), ":", 2)
					if len(credentialsPair) != 2 {
						w.Header().Set("WWW-Authenticate", "Basic realm=Authorization Required")
						cfg.UnauthorizedHandler(w, r)
						return
					}

					if credentialsPair[0] != u.UserName || credentialsPair[1] != u.Password {
						w.Header().Set("WWW-Authenticate", "Basic realm=Authorization Required")
						cfg.UnauthorizedHandler(w, r)
						return
					}
				}
			}

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}

func contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

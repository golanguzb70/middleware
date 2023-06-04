package basicauth

import (
	"github.com/gin-gonic/gin"
)

// This is configuration struct of Basic Auth
type Config struct {
	// Users is list of users that have access.
	// There may be a user1 with password1 and user2 with password2 and etc.
	Users []User `json:"users"`
	// Restricted Method means that the middleware only applies for method are given.
	// For example, PUT, POST, PATCH, DELETE methods are given to this field. Middleware check password for request with these REST Methods.
	RestrictedMethods []string `json:"restricted_methods"`
	// Restricted urls are the urls that are authoriztion is required.
	// For example, /v1/user, /v1/user/{key}, /v1/admin
	// if /v1/user is given, request url is checked for equality.
	// if /v1/user/{key} is given, request url is check for tthe urls starting with /v1/user  and one other key.
	// if /v1/user/*  is given, request url is check for all the urls that starts with '/v1/user'.
	RestrictedUrls []string `json:"restricted_urls"`
}

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Auth interface {
	Middleware(c *gin.Context)
}

// turning struct into a interface
func New(conf *Config) Auth {
	return conf
}

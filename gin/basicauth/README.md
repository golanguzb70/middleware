# Gin basic auth middleware 
This is open source and ready to use basic auth middleware package that is being maintained and contributed by experienced Golang developers for gin projects.

# Why you should use this package?
Writing basic authentication middleware by yourself for your every project takes much of your valuable time and effort. We started this open source project to save your time.


# Examples
Find example source code [here](https://github.com/golanguzb70/middleware/blob/main/gin/basicauth/example.go)

## Require Authentication for all requests
To configure your middleware to require authentication from all requests use the code below.
Here `RequireAuthForAll` field of config is set to true.

```
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/middleware/gin/basicauth"
)

func main() {
    router := RestrictAllRouter()
    router.Run(":8000")
}

func RestrictAllRouter() *gin.Engine {
	router := gin.Default()

	// This configuration checks for all incoming requests for authentication
	cfg := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: "UserName1",
				Password: "Password1",
			},
		},
		RequireAuthForAll: true,
	}

	router.Use(cfg.Middleware)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me.",
		})
	})

	router.GET("/hi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me.",
		})
	})

	return router
}
```


## Require Authentication for request with specific methods.
In some projects you need to require authentication for POST, PUT, DELETE, PATCH methods while giving access GET methods without any authorization or authentication.
This feature is best option to do so. In the example below, only request with POST, PUT, DELETE methods are require to have Authorization. 

```
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/middleware/gin/basicauth"
)

func main() {
    router := RestrictByMethodRouter()
    router.Run(":8000")
}

func RestrictByMethodRouter() *gin.Engine {
	router := gin.Default()

	cfg := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: "UserName1",
				Password: "Password1",
			},
		},
		RestrictedMethods: []string{"POST", "PUT", "DELETE"},
	}

	router.Use(cfg.Middleware)

	router.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me. Because I am POST METHOD",
		})
	})

	router.PUT("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me, because I am PUT METHOD",
		})
	})

	router.DELETE("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me, because I am DELETE  METHOD",
		})
	})

	router.GET("/user/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You can get user with id " + ctx.Param("id") + " without any authentication, because I am GET method which not restricted.",
		})
	})
	return router
}
```

## Require Authentication for request with specific urls.
In some projects, there may be a case that only some requests with GET method should require Authorization. 
For example, /admin/list should require authorization while /user/list should not. In this case, source code below helps you.

```
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/middleware/gin/basicauth"
)

func main() {
    router := RestrictByUrlRouter()
    router.Run(":8000")
}

func RestrictByUrlRouter() *gin.Engine {
	router := gin.Default()

	cfg := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: "UserName1",
				Password: "Password1",
			},
		},
		RestrictedUrls: []string{"/user/create", "/user/{id}", "/admin/*"},
	}

	router.Use(cfg.Middleware)

	router.POST("/user/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me. Because I am restricted url: /user/create.",
		})
	})

	router.DELETE("/user/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me, because I am restricted url: /user/{id}.",
		})
	})

	router.GET("/user/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me, because I am restricted url: /user/{id}.",
		})
	})

	router.POST("/admin/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me. Because I am restricted url: /admin/*.",
		})
	})

	router.DELETE("/admin/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me, because I am restricted url: /admin/*.",
		})
	})

	router.GET("/admin/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You have been asked an authentication to see me, because I am restricted url: /admin/*.",
		})
	})

	router.GET("/openurl", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You gen call me without any authentication because I am not restricted url.",
		})
	})

	return router
}

```
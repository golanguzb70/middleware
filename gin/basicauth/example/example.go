package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/middleware/gin/basicauth"
)

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
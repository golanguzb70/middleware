# Gorilla basic auth middleware 
This is open source and ready to use basic auth middleware package that is being maintained and contributed by experienced Golang developers for gorilla projects.

# Why you should use this package?
Writing basic auth middleware by yourself for your every project takes much of your valuable time and effort. We started this open source project to save your time.


# Examples
Find example source code [here](https://github.com/golanguzb70/middleware/blob/main/gorilla/basicauth/example.go)

## Require Authentication for all requests
To configure your middleware to require authentication from all requests use the code below.
Here `RequireAuthForAll` field of config is set to true.
```go
package main

import (
	"log"
	"net/http"

	"github.com/golanguzb70/middleware/gorilla/basicauth"
	"github.com/gorilla/mux"
)

func main() {
	r := RequireAuthForAll()
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

func RequireAuthForAll() *mux.Router {
	router := mux.NewRouter()

	// This configuration checks for all incoming requests for authentication
	cfg := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: "username",
				Password: "password",
			},
			{
				UserName: "username2",
				Password: "password2",
			},
		},
		RequireAuthForAll: true,
		UnauthorizedHandler: func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		},
	}
	router.Use(basicauth.Middleware(cfg))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me."}`))
	}).Methods("GET")

	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me."}`))
	}).Methods("GET")

	return router
}
```



## Require Authentication for request with specific methods.
In some projects you need to require authentication for POST, PUT, DELETE, PATCH methods while giving access GET methods without any authorization or authentication.
This feature is best option to do so. In the example below, only request with POST, PUT, DELETE methods are require to have Authorization. 
```go
package main

import (
	"log"
	"net/http"

	"github.com/golanguzb70/middleware/gorilla/basicauth"
	"github.com/gorilla/mux"
)

func main() {
	r := RestrictByMethodsRouter()
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

func RestrictByMethodsRouter() *mux.Router {
	router := mux.NewRouter()

	// This configuration checks for all incoming requests for authentication
	cfg := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: "username",
				Password: "password",
			},
			{
				UserName: "username2",
				Password: "password2",
			},
		},
		RestrictedMethods: []string{"POST", "DELETE", "PUT"},
		UnauthorizedHandler: func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		},
	}
	router.Use(basicauth.Middleware(cfg))

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me. Because I am POST METHOD"}`))
	}).Methods("POST")

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me, because I am PUT METHOD"}`))
	}).Methods("PUT")

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me, because I am DELETE  METHOD"}`))
	}).Methods("DELETE")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte(`{"message": "You can get user with id "` + vars["id"] + `" without any authentication, because I am GET method which not restricted."}`))
	}).Methods("GET")

	return router
}
```



## Require Authentication for request with specific urls.
In some projects, there may be a case that only some requests with GET method should require Authorization. 
For example, /admin/list should require authorization while /user/list should not. In this case, source code below helps you.

```go
package main

import (
	"log"
	"net/http"

	"github.com/golanguzb70/middleware/gorilla/basicauth"
	"github.com/gorilla/mux"
)

func main() {
	r := RestrictByUrlRouter()
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

func RestrictByUrlRouter() *mux.Router {
	router := mux.NewRouter()

	// This configuration checks for all incoming requests for authentication
	cfg := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: "username",
				Password: "password",
			},
			{
				UserName: "username2",
				Password: "password2",
			},
		},
		RestrictedUrls: []string{"/user/create", "/user/{id}", "/admin/*"},
		UnauthorizedHandler: func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		},
	}
	router.Use(basicauth.Middleware(cfg))

	router.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me. Because I am restricted url: /user/create."}`))
	}).Methods("POST")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me, because I am restricted url: /user/{id}."}`))
	}).Methods("GET")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me, because I am restricted url: /user/{id}."}`))
	}).Methods("DELETE")

	router.HandleFunc("/admin/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me. Because I am restricted url: /admin/*."}`))
	}).Methods("POST")

	router.HandleFunc("/admin/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You have been asked an authentication to see me, because I am restricted url: /admin/*."}`))
	}).Methods("GET")

	router.HandleFunc("/openurl", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You gen call me without any authentication because I am not restricted url."}`))
	}).Methods("GET")

	return router
}
```
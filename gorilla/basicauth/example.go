package basicauth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RestrictAllRouter() *mux.Router {
	router := mux.NewRouter()

	// Create a config with your desired values
	config := Config{
		Users: []User{
			{
				UserName: "user1",
				Password: "password1",
			},
		},
		RequireAuthForAll: true,
		UnauthorizedHandler: func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		},
	}

	// Apply the Basic Auth middleware to the router
	router.Use(Middleware(config))

	// Define your routes and handlers
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me."))
	}).Methods("GET")

	router.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me."))
	}).Methods("GET")

	return router
}

func RestrictByUrlRouter() *mux.Router {
	router := mux.NewRouter()
	config := Config{
		Users: []User{
			{
				UserName: "user1",
				Password: "password1",
			},
		},
		RestrictedUrls: []string{"/user/create", "/admin/*", "user/{id}"},
		UnauthorizedHandler: func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		},
	}
	// Apply the Basic Auth middleware to the router
	router.Use(Middleware(config))

	router.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me. Because I am restricted url: /user/create"))
	}).Methods("POST")

	router.HandleFunc("/user/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me. Because I am restricted url: /user/{id}"))
	}).Methods("DELETE")

	router.HandleFunc("/user/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me. Because I am restricted url: /user/{id}"))
	}).Methods("GET")

	router.HandleFunc("/admin/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me. Because I am restricted url: /admin/*"))
	}).Methods("POST")

	router.HandleFunc("/admin/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me. Because I am restricted url: /admin/*"))
	}).Methods("GET")

	router.HandleFunc("/admin/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me. Because I am restricted url: /admin/*"))
	}).Methods("DELETE")

	router.HandleFunc("/openurl", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You gen call me without any authentication because I am not restricted url."))
	}).Methods("GET")

	return router
}

func RestrictByMethodsRouter() *mux.Router {
	router := mux.NewRouter()
	config := Config{
		Users: []User{
			{
				UserName: "user",
				Password: "password",
			},
		},
		RestrictedMethods: []string{"POST", "GET", "DELETE", "PUT"},
		UnauthorizedHandler: func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		},
	}

	// Apply the Basic Auth middleware to the router
	router.Use(Middleware(config))

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me. Because I am POST METHOD"))
	}).Methods("POST")

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me, because I am PUT METHOD"))
	}).Methods("PUT")

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have been asked an authentication to see me, because I am DELETE METHOD"))
	}).Methods("DELETE")

	router.HandleFunc("/user/:id", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("You can get user with id " + vars["id"] + " without any authentication, because I am GET method which not restricted."))
	}).Methods("GET")

	return router
}

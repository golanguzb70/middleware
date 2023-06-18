package basicauth

import (
	"encoding/base64"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestRequireAuthorizationAll(t *testing.T) {
	cfg := Config{
		Users: []User{
			{
				UserName: "user",
				Password: "password",
			},
		},
		RequireAuthForAll: true,
	}
	// Create a new router
	router := RestrictAllRouter()

	// Create a test request without authorization header
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/hi", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Result().StatusCode)

	// Create a test request with valid authorization header
	base64Auth := base64.StdEncoding.EncodeToString([]byte(cfg.Users[0].UserName + ":" + cfg.Users[0].Password))

	reqWithAuth := httptest.NewRequest("GET", "/", nil)
	reqWithAuth.Header.Set("Authorization", "Basic "+base64Auth)
	wWithAuth := httptest.NewRecorder()
	router.ServeHTTP(wWithAuth, reqWithAuth)

	// assert.Equal(t, 200, wWithAuth.Result().StatusCode)
}
func TestRequireForSpecificMethods(t *testing.T) {
	cfg := Config{
		Users: []User{
			{
				UserName: "username",
				Password: "password",
			},
		},
	}

	// Create a new router
	router := RestrictByMethodsRouter()

	// Create a test request without authorization header
	req := httptest.NewRequest("POST", "/user", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("PUT", "/user", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("DELETE", "/user", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/user/10", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)

	// Create a test request with valid authorization header
	base64Auth := base64.StdEncoding.EncodeToString([]byte(cfg.Users[0].UserName + ":" + cfg.Users[0].Password))

	req = httptest.NewRequest("POST", "/user", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("PUT", "/user", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("DELETE", "/user", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/user/10", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)
}

func TestRequireForSpecificUrls(t *testing.T) {
	cfg := Config{
		Users: []User{
			{
				UserName: "username",
				Password: "password",
			},
		},
	}
	// Set Gin to Test mode

	// Create a new router
	router := RestrictByUrlRouter()
	// Create a test request without authorization header
	req := httptest.NewRequest("POST", "/user/create", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("DELETE", "/user/12", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/user/12", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("POST", "/admin/create", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("DELETE", "/admin/10", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/admin/10", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/openurl", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	// Create a test request with authorization header
	base64Auth := base64.StdEncoding.EncodeToString([]byte(cfg.Users[0].UserName + ":" + cfg.Users[0].Password))

	req = httptest.NewRequest("POST", "/user/create", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("DELETE", "/user/12", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/user/12", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("POST", "/admin/create", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("DELETE", "/admin/10", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/admin/10", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)

	req = httptest.NewRequest("GET", "/openurl", nil)
	req.Header.Set("Authorization", "Basic "+base64Auth)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestHelloWorld(t *testing.T) {
	body := gin.H{
		"hello": "world",
	}

	router := SetupRouter()

	w := performRequest(router, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["hello"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["hello"], value)
}

func TestPlus(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "POST", "/plus/9/7")

	var response map[string]int
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["result"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, value, 16)
}

func TestMultiple(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "POST", "/multiple/9/7")

	result, _ := strconv.ParseFloat(w.Body.String(), 64)

	assert.Equal(t, result, float64(9*7))
}

func TestProfile(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "GET", "/profile")

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	status, _ := response["status"]
	data, _ := response["data"].(map[string]interface{})

	assert.Nil(t, err)
	assert.Equal(t, status, float64(200))
	assert.Equal(t, data["name"], "Surya")
	assert.Equal(t, data["role"], "Developer")
}

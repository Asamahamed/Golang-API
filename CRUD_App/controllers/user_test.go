package controllers

import (
	"Crud_Api/config"
	"Crud_Api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	config.Connect()
	config.DB.AutoMigrate(&models.User{})

	payload := map[string]string{
		"name":  "Asam",
		"email": "asam@gmail.com",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users", CreateUser).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201")
	var responseUser models.User
	json.Unmarshal(rr.Body.Bytes(), &responseUser)
	assert.Equal(t, "Asam", responseUser.Name, "User name should match")
	assert.Equal(t, "asam@gmail.com", responseUser.Email, "User email should match")
}

func TestGetUsers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users", GetUsers).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	var users []models.User
	json.Unmarshal(rr.Body.Bytes(), &users)
	assert.NotEmpty(t, users, "Users should not be empty")
}

func TestUpdateUser(t *testing.T) {

	config.Connect()
	config.DB.AutoMigrate(&models.User{})

	user := models.User{Name: "Asam", Email: "asam@gmail.com"}
	config.DB.Create(&user)

	payload := map[string]string{
		"name":  "Asam Ahamed",
		"email": "asam_Ahamed@gmail.com",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", "/api/v1/users/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users/{id}", UpdateUser).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	var updatedUser models.User
	json.Unmarshal(rr.Body.Bytes(), &updatedUser)
	assert.Equal(t, "Asam Ahamed", updatedUser.Name, "User name should match")
	assert.Equal(t, "asam_Ahamed@gmail.com", updatedUser.Email, "User email should match")
}

func TestDeleteUser(t *testing.T) {

	config.Connect()
	config.DB.AutoMigrate(&models.User{})

	user := models.User{Name: "Asam", Email: "asam@gmail.com"}
	config.DB.Create(&user)

	req, _ := http.NewRequest("DELETE", "/api/v1/users/"+strconv.Itoa(int(user.ID)), nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users/{id}", DeleteUser).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.Equal(t, `{"message":"User deleted successfully"}`, rr.Body.String())
}

package routes

import (
	"Crud_Api/controllers"

	"github.com/gorilla/mux"
)

func InitializeUserRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", controllers.DeleteUser).Methods("DELETE")

}

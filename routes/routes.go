package routes

import (
	"go-mongodb/app/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/project", controllers.GetProjectListEndpoint).Methods("GET")
	router.HandleFunc("/project", controllers.CreateProjectEndpoint).Methods("POST")
	router.HandleFunc("/project/{id}", controllers.GetProjectEndpoint).Methods("GET")
	router.HandleFunc("/project/{id}", controllers.DeleteProjectEndpoint).Methods("DELETE")
	router.HandleFunc("/project/{id}", controllers.UpdateProjectEndpoint).Methods("PUT")

	return router
}

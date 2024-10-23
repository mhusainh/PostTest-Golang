package routes

import (
	"Weekly-Task3/pkg/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")

	return router
}

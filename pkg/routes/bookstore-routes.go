package routes

import (
	"github.com/gorilla/mux"
	"github.com/Shadabdullah/GoBook/pkg/controllers"

)

var RegisterBookStoreRoutes = func(router *mux.Router)  {
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}",UpdateBook).Methods("PUT")

	

	
	
}
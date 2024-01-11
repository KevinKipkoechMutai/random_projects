package routes

import (
	"github.com/gorilla/mux"
	"github.com/KevinKipkoechMutai/random_projects/tree/main/book-management-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{book_id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods("DELETE")
}
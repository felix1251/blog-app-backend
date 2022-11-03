package routes

import (
	"github.com/gorrila/mux"
	"github.com/felix1251/go-rest-api/pkg/controllers"
)

var RegisterPostRoutes = func(router *mux.Router){
	router.HandleFunc("/post/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post/", controllers.GetPost).Methods("GET")
	router.HandleFunc("/post/{postId}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/post/{postId}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{postId}", controllers.DeletePost).Methods("DELETE")
}

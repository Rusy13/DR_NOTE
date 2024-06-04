package routes

import (
	"awesomeProject/internal/middleware"
	"awesomeProject/internal/user/delivery"
	"github.com/gorilla/mux"
)

func GetRouter(handlers *delivery.UserDelivery, mw *middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	assignRoutes(router, handlers)
	assignMiddleware(router, mw)
	return router
}

func assignRoutes(router *mux.Router, handlers *delivery.UserDelivery) {
	router.HandleFunc("/users", handlers.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}/subscribers", handlers.GetSubscribers).Methods("GET")
	router.HandleFunc("/users/{userID}/subscribe/{subscribedToID}", handlers.Subscribe).Methods("POST")
	router.HandleFunc("/users/{userID}/unsubscribe/{subscribedToID}", handlers.Unsubscribe).Methods("POST")

}

func assignMiddleware(router *mux.Router, mw *middleware.Middleware) {
	router.Use(mw.AccessLog)
	//router.Use(mw.Auth)
}

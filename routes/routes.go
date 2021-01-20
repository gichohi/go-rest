package routes

import (
	"github.com/gichohi/go-rest.git/controllers"
	"github.com/gichohi/go-rest.git/utils"
	"net/http"
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(CommonMiddleware)
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	// Auth route
	s := router.PathPrefix("/auth").Subrouter()
	s.Use(utils.JwtVerify)
	s.HandleFunc("/user/{username}", controllers.GetUser).Methods("GET")
	return router
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

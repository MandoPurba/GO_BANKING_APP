package apps

import (
	"github.com/MandoPurba/rest-api/apps/controller"
	err "github.com/MandoPurba/rest-api/utils/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() http.Handler {
	router := mux.NewRouter()
	// APP ROUTE

	// USER ROUTER
	router.HandleFunc("/users", controller.GetALlUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/create-user", controller.CreateUser).Methods("POST")

	// ACCOUNT ROUTER
	router.HandleFunc("/activate-account", controller.ActivateAccount).Methods("POST")

	// TRANSFER ROUTER
	router.HandleFunc("/transfer", controller.Transfer).Methods("POST")
	// ----------------- //

	router.NotFoundHandler = http.HandlerFunc(err.NotFoundHandler)
	return router
}

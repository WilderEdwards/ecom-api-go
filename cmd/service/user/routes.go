//where API endpoints are handled

package user

import (
	"net/http"
	"github.com/gorilla/mux"
)
type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	//handling POST requests for both user login and register
	router.HandleFunc("/login", h.handleLogin).Methods("POST") 
	router.HandleFunc("/register", h.handleLogin).Methods("POST") 

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//function to handle api request for login
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//function to handle api request for *register*, similair to maintaian modularity
}

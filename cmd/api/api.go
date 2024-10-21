package api

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/WilderEdwards/ecom-api-go/cmd/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	//add different features to this function for the API ; modular
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()



	//example of feature addition
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	//prefixed with v1.api via subrouter ^

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

package server

import (
	"CashTest/customer"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(cst customer.CustomerService) *Server {
	s := Server{router: mux.NewRouter()}
	NewCustomerRouter(cst, s.newSubrouter("/customer"))
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port 3000")
	if err := http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) newSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}

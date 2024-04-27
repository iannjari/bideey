package api

import (
	"bideey/handler"
	"bideey/util"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	address string
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", sayHello())
	router.HandleFunc("POST /bid", handler.CreateBid())
	router.HandleFunc("PUT /bid", handler.UpdateBid())
	router.HandleFunc("DELETE /bid/{id}", handler.DeleteBid())

	server := http.Server{
		Addr:    s.address,
		Handler: router,
	}

	log.Println("Server started, listening on:", s.address)

	return server.ListenAndServe()
}

func NewServer(hostAdress string) (server *Server) {
	server = &Server{
		address: hostAdress,
	}
	return
}

func sayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(util.GetRandomStrCode())
	}
}

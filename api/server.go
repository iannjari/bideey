package api

import (
	"bideey/handler"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	address string
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	// hello
	router.HandleFunc("GET /hello", sayHello())
	router.HandleFunc("GET /", sayHello())

	// bids controllers
	router.HandleFunc("POST /bid", handler.CreateBid())
	router.HandleFunc("PUT /bid", handler.UpdateBid())
	router.HandleFunc("DELETE /bid/{id}", handler.DeleteBid())

	// biddables controllers
	router.HandleFunc("POST /biddable", handler.CreateBiddable())
	router.HandleFunc("PUT /biddable", handler.UpdateBiddable())
	router.HandleFunc("DELETE /biddable/{id}", handler.DeleteBiddable())

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

		json.NewEncoder(w).Encode("Hello you! 😊")
	}
}

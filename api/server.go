package api

import (
	"bideey/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	address int
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", sayHello())

	addr := ":" + fmt.Sprint(s.address)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("Server started, listening on port:", s.address)

	return server.ListenAndServe()
}

func NewServer(address int) (server *Server) {
	server = &Server{
		address: address,
	}
	return
}

func sayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(util.GetRandomStrCode())
	}
}

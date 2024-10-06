package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"Hi_Tech/internal/controller"
	"Hi_Tech/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	p := &controller.Product{}

	r.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		err := p.GetAllProducts(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	r.HandleFunc("/get/{id}", func(writer http.ResponseWriter, request *http.Request) {
		err := p.GetProductById(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	return r
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	s := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

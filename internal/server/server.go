package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	u := &controller.User{}

	r.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		err := p.GetAllProducts(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	r.HandleFunc("/get/{id}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		err := p.GetProductById(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	r.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		err := u.RegisterUser(writer, request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		} else {
			writer.WriteHeader(http.StatusCreated)
		}
	})
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		jsonResponse, err := json.Marshal(map[string]string{"message": "Hello, World!"})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsonResponse)
	})

	return r
}

func NewServer() *http.Server {
	port := 8080
	s := &Server{
		port: port,
		db:   database.DBConnection(),
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

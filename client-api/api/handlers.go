package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type TestResponse struct {
			Test string `json:"test"`
		}
		testResponse := TestResponse{Test: "alive"}
		testJSON, err := json.Marshal(testResponse)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err = w.Write(testJSON); err != nil {
			log.Printf("Error handling health check: %s", err)
		}
	}
}

func (s *Server) handleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ports, err := s.Client.List(context.Background())
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error calling grpc..."))
			w.Write([]byte(err.Error()))
			return
		}
		resp, err := json.Marshal(ports)
		if err != nil {
		}
		w.WriteHeader(200)
		w.Write(resp)
	}
}

func (s *Server) handleSinglePort() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, ok := mux.Vars(r)["code"]
		if !ok {
			log.Println("Not valid Port CODE provided (should be XXXXX).")
			w.WriteHeader(400)
			w.Write([]byte("Invalid code."))
			return
		}
		ports, err := s.Client.GetByID(context.Background(), code)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error calling grpc..."))
			return
		}
		resp, err := json.Marshal(ports)
		if err != nil {
		}
		w.WriteHeader(200)
		w.Write(resp)
	}
}

func (s *Server) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, ok := mux.Vars(r)["code"]
		if !ok {
			log.Println("Not valid Port CODE provided (should be XXXXX).")
			w.WriteHeader(400)
			w.Write([]byte("Invalid code."))
			return
		}
		// Better error-handling? Split into datalayer errors and transport errors!
		ports, err := s.Client.DeleteByID(context.Background(), code)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Error calling grpc..."))
			return
		}
		resp, err := json.Marshal(ports)
		if err != nil {
		}
		w.WriteHeader(200)
		w.Write(resp)
	}
}

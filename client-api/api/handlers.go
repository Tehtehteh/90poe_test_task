package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"t_task/proto"
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
		ports, err := s.PDServiceClient.List(context.Background(), &proto.NoParams{})
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

package api

func (s *Server) routes() {
	s.Router.HandleFunc("/ping", s.handleHealthCheck())
	s.Router.HandleFunc("/ports", s.handleList())
	s.Router.HandleFunc("/ports/{code}", s.handleSinglePort())
	s.Router.HandleFunc("/ports/{code}/delete", s.handleDelete())
}

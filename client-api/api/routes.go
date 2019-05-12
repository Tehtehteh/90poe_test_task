package api

func (s *Server) routes() {
	s.Router.HandleFunc("/ping", s.handleHealthCheck())
	s.Router.HandleFunc("/parser", s.handleList())
	s.Router.HandleFunc("/parser/{code}", s.handleSinglePort())
	s.Router.HandleFunc("/parser/{code}/delete", s.handleDelete())
}

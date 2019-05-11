package api

func (s *Server) routes() {
	s.Router.HandleFunc("/ping", s.handleHealthCheck())
	s.Router.HandleFunc("/ports", s.handleList())
	s.Router.HandleFunc("/ports/{code}", s.handleList())
	//s.Router.HandleFunc("/{domain}/videos", s.videosHandler())
	//s.Router.HandleFunc("/{domain}/video-aggregate", s.videoAggregateHandler())
	//s.Router.HandleFunc("/{domain}/{userID}/user", s.userHandler())
}

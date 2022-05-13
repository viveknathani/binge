package server

func (s *Server) SetupRoutes() {

	s.Router.HandleFunc("/", s.serveHomePage)
	s.Router.HandleFunc("/show", s.serveShowEpisodesPage)
	s.Router.HandleFunc("/watch", s.serveWatchVideoPage)
	s.Router.HandleFunc("/upload", s.serveUploadPage)
	s.Router.HandleFunc("/uploadVideo", s.uploadHandler)
	s.Router.HandleFunc("/jobs", s.servePendingJobsPage)
	s.Router.Use(setContentTypeFileFormat)
	s.setupContent("content")
}

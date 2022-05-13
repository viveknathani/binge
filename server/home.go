package server

import (
	"html/template"
	"net/http"

	"github.com/viveknathani/binge/entity"
)

type homePageVariables struct {
	ShowsList  *[]entity.Show
	MoviesList *[]entity.Movie
}

func (s *Server) serveHomePage(w http.ResponseWriter, r *http.Request) {

	shows, movies, err := s.Service.GetShowsAndMovies(r.Context())
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}

	t, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}
	err = t.Execute(w, homePageVariables{
		ShowsList:  shows,
		MoviesList: movies,
	})
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}
}

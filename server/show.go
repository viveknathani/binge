package server

import (
	"html/template"
	"net/http"

	"github.com/viveknathani/binge/entity"
)

type showPageVariables struct {
	Episodes *[]entity.Episode
}

func (s *Server) serveShowEpisodesPage(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	episodes, err := s.Service.GetEpisodes(r.Context(), params["s"][0])
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}

	t, err := template.ParseFiles("web/templates/show.html")
	err = t.Execute(w, showPageVariables{
		Episodes: episodes,
	})
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}
}

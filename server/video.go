package server

import (
	"html/template"
	"net/http"
)

type watchPageVariables struct {
	VideoId string
}

func (s *Server) serveWatchVideoPage(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	videoId := params["v"][0]
	t, err := template.ParseFiles("web/templates/video.html")
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}

	err = t.Execute(w, watchPageVariables{
		VideoId: videoId,
	})
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}
}

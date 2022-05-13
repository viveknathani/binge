package server

import (
	"html/template"
	"net/http"
)

type jobsPageVariables struct {
	JobsList []string
}

func (s *Server) servePendingJobsPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("web/templates/jobs.html")
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}

	err = t.Execute(w, jobsPageVariables{
		JobsList: s.Service.Processor.GetAllPendingJobs(),
	})
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}
}

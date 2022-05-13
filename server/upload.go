package server

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/viveknathani/binge/entity"
	"github.com/viveknathani/binge/processor"
)

func (s *Server) serveUploadPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("web/templates/upload.html")
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		if ok := sendServerError(w); ok != nil {
			s.Service.Logger.Error(err.Error(), zapReqID(r))
		}
		return
	}
}

func (s *Server) uploadHandler(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("file")
	movieName := r.FormValue("name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	videoID := time.Now().UnixNano()
	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", videoID, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := &entity.Video{
		VideoId: fmt.Sprintf("%d", videoID),
	}
	s.Service.AddVideo(r.Context(), v)

	m := &entity.Movie{
		Name:    movieName,
		VideoId: v.VideoId,
	}
	s.Service.AddMovie(r.Context(), m)

	s.Service.Processor.Push(processor.Event{
		Path:    fmt.Sprintf("./uploads/%d%s", videoID, filepath.Ext(fileHeader.Filename)),
		VideoId: v.VideoId,
	})

	fmt.Fprintf(w, "Upload successful")
}

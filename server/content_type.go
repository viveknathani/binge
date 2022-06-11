package server

import (
	"net/http"
	"strings"
)

func hasExtension(path string, extension string) bool {
	return strings.HasSuffix(path, extension) || strings.HasSuffix(path, extension+"/")
}

func setContentTypeFileFormat(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		value := "text/html"

		if hasExtension(r.URL.Path, ".css") {
			value = "text/css"
		}

		if hasExtension(r.URL.Path, ".js") {
			value = "text/javascript"
		}

		if hasExtension(r.URL.Path, ".mpd") {
			value = "application/dash+xml"
		}

		if hasExtension(r.URL.Path, ".m4s") {
			value = "video/mp4"
		}

		w.Header().Add("Content-Type", value)
		handler.ServeHTTP(w, r)
	})
}

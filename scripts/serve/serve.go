package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrefrco/resume/scripts/resume"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, err := resume.RenderResumeHTML()
		if err != nil {
			http.Error(w, "Failed to render resume: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if html == nil {
			http.Error(w, "Resume data not available", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	})

	port := 8080
	log.Printf("Serving on http://localhost:%d\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

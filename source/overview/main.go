package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type MicroserviceStatus struct {
	FrontendActiveVersion  string
	FrontendPreviewVersion string
	FrontendInPreview      bool
	BackendActiveVersion   string
	BackendPreviewVersion  string
	BackendInPreview       bool
	WorkerActiveVersion    string
	WorkerPreviewVersion   string
	WorkerInPreview        bool
	QueueActiveName        string
	QueuePreviewName       string
	QueueInPreview         bool

	BackendVersion string
	BackendHost    string
	BackendPort    string
	LoanAmount     int
	LoanResult     string
}

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	microserviceStatus := MicroserviceStatus{}

	// Kubernetes check if app is ok
	http.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "up")
	})

	// Kubernetes check if app can serve requests
	http.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "yes")
	})

	http.HandleFunc("/", microserviceStatus.serveFiles)

	fmt.Printf("Microservice overview is now at port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

func (microserviceStatus *MicroserviceStatus) serveFiles(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	p := "." + upath
	if p == "./" {
		microserviceStatus.home(w, r)
		return
	} else if p == "./diagram.svg" {
		microserviceStatus.renderLiveDiagram(w, r)
		return
	} else {
		p = filepath.Join("./static/", path.Clean(upath))
	}
	http.ServeFile(w, r, p)
}

func (microserviceStatus *MicroserviceStatus) home(w http.ResponseWriter, r *http.Request) {

	// microserviceStatus.findBackendVersion()
	microserviceStatus.findStatus()

	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}
	err = t.Execute(w, microserviceStatus)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}

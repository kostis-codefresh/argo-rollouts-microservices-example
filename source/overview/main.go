package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type MicroserviceStatus struct {
	AppVersion     string
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

	microserviceStatus.AppVersion = os.Getenv("APP_VERSION")
	if len(microserviceStatus.AppVersion) == 0 {
		microserviceStatus.AppVersion = "dev"
	}

	microserviceStatus.BackendHost = os.Getenv("BACKEND_HOST")
	if len(microserviceStatus.BackendHost) == 0 {
		microserviceStatus.BackendHost = "interest"
	}

	microserviceStatus.BackendPort = os.Getenv("BACKEND_PORT")
	if len(microserviceStatus.BackendPort) == 0 {
		microserviceStatus.BackendPort = "8080"
	}

	// Allow anybody to retrieve version
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, microserviceStatus.AppVersion)
	})

	// Kubernetes check if app is ok
	http.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "up")
	})

	// Kubernetes check if app can serve requests
	http.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "yes")
	})

	http.HandleFunc("/", microserviceStatus.serveFiles)

	fmt.Printf("Frontend version %s is listening now at port %s\n", microserviceStatus.AppVersion, port)
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

func (microserviceStatus *MicroserviceStatus) findBackendVersion() {
	version, err := microserviceStatus.callBackend("version")
	if err != nil {
		log.Println("Interest error :", err)
		version = "unknown"
	}

	microserviceStatus.BackendVersion = version
}

func (microserviceStatus *MicroserviceStatus) home(w http.ResponseWriter, r *http.Request) {

	microserviceStatus.findBackendVersion()

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

func (microserviceStatus *MicroserviceStatus) callBackend(path string) (result string, err error) {

	backendUrl := url.URL{
		Scheme: "http",
		Host:   microserviceStatus.BackendHost + ":" + microserviceStatus.BackendPort,
		Path:   path,
	}

	url := backendUrl.String()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Could not access %s, got %s\n ", url, err)
		return "", errors.New("Could not access " + url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
		return "", errors.New("Could not access " + url)
	}

	log.Printf("Response status of %s: %s\n", url, resp.Status)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

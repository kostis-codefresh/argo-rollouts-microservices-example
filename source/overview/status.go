package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/url"
)

func (microserviceStatus *MicroserviceStatus) findStatus() {

	microserviceStatus.FrontendActiveVersion = "1.0"
	microserviceStatus.FrontendPreviewVersion = "2.0"
	microserviceStatus.FrontendInPreview = true
	microserviceStatus.BackendActiveVersion = "1.0"
	microserviceStatus.BackendPreviewVersion = "2.0"
	microserviceStatus.BackendInPreview = true
	microserviceStatus.WorkerActiveVersion = "1.0"
	microserviceStatus.WorkerPreviewVersion = "2.0"
	microserviceStatus.WorkerInPreview = true
	microserviceStatus.QueueActiveName = "production queue"
	microserviceStatus.QueuePreviewName = "preview queue"
	microserviceStatus.QueueInPreview = true
}

func (microserviceStatus *MicroserviceStatus) findBackendVersion() {
	version, err := microserviceStatus.callBackend("version")
	if err != nil {
		log.Println("Interest error :", err)
		version = "unknown"
	}

	microserviceStatus.BackendVersion = version
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

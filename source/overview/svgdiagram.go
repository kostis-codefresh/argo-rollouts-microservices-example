package main

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
)

const (
	sceneWidth  = 1200
	sceneHeight = 800

	rolloutWidth  = 200
	rolloutHeight = 100

	queueDiskHeight = 20
	queueDiskWidth  = 50
)

func (microserviceStatus *MicroserviceStatus) renderLiveDiagram(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(sceneWidth, sceneHeight)
	s.Circle(600, 600, 125, `fill:none;stroke:black`, `id="circle"`)
	s.Animate("#circle", "opacity", 0, 1, 3, 15)

	s.Rect(10, 200, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="frontendActive"`)

	s.Rect(300, 200, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	s.Rect(500, 30, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="workerActive"`)

	s.Ellipse(900, 220, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 200, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 180, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 160, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 140, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 120, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 100, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	s.Line(250, 240, 280, 250, `stroke-width:3;stroke:black`)
	s.Line(220, 250, 280, 250, `stroke-width:3;stroke:black`)
	s.Line(250, 260, 280, 250, `stroke-width:3;stroke:black`)

	s.End()
}

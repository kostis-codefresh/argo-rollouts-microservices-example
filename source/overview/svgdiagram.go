package main

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func (microserviceStatus *MicroserviceStatus) renderLiveDiagram(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(1200, 800)
	s.Circle(600, 600, 125, `fill:none;stroke:black`, `id="circle"`)
	s.Animate("#circle", "opacity", 0, 1, 3, 15)

	s.Rect(10, 200, 200, 100, `fill:lightblue;stroke:black`, `id="frontendActive"`)

	s.Rect(300, 200, 200, 100, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	s.Rect(500, 30, 200, 100, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	s.Ellipse(900, 220, 50, 20, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 200, 50, 20, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 180, 50, 20, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 160, 50, 20, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 140, 50, 20, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 120, 50, 20, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	s.Ellipse(900, 100, 50, 20, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	s.End()
}

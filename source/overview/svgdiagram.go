package main

import (
	"fmt"
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

	numberOfDisks = 7
	diskSpacing   = 20

	margin = 10
)

type queueGFX struct {
	color       string
	name        string
	description string
}

func (microserviceStatus *MicroserviceStatus) renderLiveDiagram(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(sceneWidth, sceneHeight)
	s.Circle(600, 600, 125, `fill:none;stroke:black`, `id="circle"`)
	s.Animate("#circle", "opacity", 0, 1, 3, 15)

	s.Rect(10, 200, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="frontendActive"`)

	s.Rect(300, 200, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	s.Rect(500, 30, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="workerActive"`)

	// s.Ellipse(900, 220, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	// s.Ellipse(900, 200, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	// s.Ellipse(900, 180, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	// s.Ellipse(900, 160, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	// s.Ellipse(900, 140, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	// s.Ellipse(900, 120, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)
	// s.Ellipse(900, 100, queueDiskWidth, queueDiskHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	q := queueGFX{}
	q.color = "orange"
	q.name = "production"
	q.description = "RabbitMQ"
	microserviceStatus.renderQueue(s, q)

	s.Line(250, 240, 280, 250, `stroke-width:3;stroke:black`)
	s.Line(220, 250, 280, 250, `stroke-width:3;stroke:black`)
	s.Line(250, 260, 280, 250, `stroke-width:3;stroke:black`)

	s.Text(10, 200, "Front-end", "font-size:30px;fill:black")
	s.Text(300, 200, "Back-end", "font-size:30px;fill:black")
	s.Text(500, 30, "Worker", "font-size:30px;fill:black")

	s.End()
}

func (microserviceStatus *MicroserviceStatus) renderQueue(canvas *svg.SVG, queue queueGFX) {
	x := 900
	y := 100

	fontSize := 26
	svgTextOptions := fmt.Sprintf("font-size:%dpx;fill:black", fontSize)
	svgDiskOptions := fmt.Sprintf("fill:%s;stroke:black", queue.color)

	startingY := y + (diskSpacing * numberOfDisks)
	for i := 0; i < numberOfDisks; i++ {
		canvas.Ellipse(x, startingY, queueDiskWidth, queueDiskHeight, svgDiskOptions, `id="frontendPreview"`)
		startingY = startingY - diskSpacing
	}

	//Description on top of the disks
	canvas.Text(x-queueDiskWidth, y-margin, queue.description, svgTextOptions)

	//queue name is to the right of the disks
	canvas.Text(x+queueDiskWidth+margin, y+(diskSpacing*numberOfDisks/2), queue.name, svgTextOptions)

}

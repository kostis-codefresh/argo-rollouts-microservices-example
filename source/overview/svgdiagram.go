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
	x           int
	y           int
}

type microserviceGFX struct {
	color   string
	version string
	name    string
	x       int
	y       int
}

type sceneGFX struct {
	microserviceStatus *MicroserviceStatus
	canvas             *svg.SVG

	frontendStable  microserviceGFX
	frontendPreview microserviceGFX
	backendStable   microserviceGFX
	backendPreview  microserviceGFX
	workerStable    microserviceGFX
	workerPreview   microserviceGFX

	queueStable  queueGFX
	queuePreview queueGFX
}

func (microserviceStatus *MicroserviceStatus) renderLiveDiagram(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	scene := sceneGFX{}
	s := svg.New(w)

	scene.canvas = s
	scene.prepareScene(microserviceStatus)

	s.Start(sceneWidth, sceneHeight)
	s.Circle(600, 400, 50, `fill:none;stroke:black`, `id="circle"`)
	s.Animate("#circle", "opacity", 0, 1, 3, 15)

	// s.Rect(10, 200, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="frontendActive"`)
	scene.renderMicroservice(scene.frontendStable)
	scene.renderMicroservice(scene.frontendPreview)

	scene.renderMicroservice(scene.backendStable)
	scene.renderMicroservice(scene.backendPreview)

	// s.Rect(300, 200, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="frontendPreview"`)

	s.Rect(500, 30, rolloutWidth, rolloutHeight, `fill:lightblue;stroke:black`, `id="workerActive"`)

	scene.renderQueue(scene.queueStable)
	scene.renderQueue(scene.queuePreview)

	s.Line(250, 240, 280, 250, `stroke-width:3;stroke:black`)
	s.Line(220, 250, 280, 250, `stroke-width:3;stroke:black`)
	s.Line(250, 260, 280, 250, `stroke-width:3;stroke:black`)

	// s.Text(10, 200, "Front-end", "font-size:30px;fill:black")
	// s.Text(300, 200, "Back-end", "font-size:30px;fill:black")
	s.Text(500, 30, "Worker", "font-size:30px;fill:black")

	s.End()
}

func (scene *sceneGFX) prepareScene(microserviceStatus *MicroserviceStatus) {
	scene.microserviceStatus = microserviceStatus

	//Production frontend
	frontendStable := microserviceGFX{}
	frontendStable.color = "blue"
	frontendStable.name = "Frontend"
	frontendStable.version = "1.0"
	frontendStable.x = 10
	frontendStable.y = 200
	scene.frontendStable = frontendStable

	//Preview frontend
	frontendPreview := microserviceGFX{}
	frontendPreview.color = "green"
	frontendPreview.name = "Frontend"
	frontendPreview.version = "2.0"
	frontendPreview.x = 10
	frontendPreview.y = 600
	scene.frontendPreview = frontendPreview

	//Production backend
	backendStable := microserviceGFX{}
	backendStable.color = "blue"
	backendStable.name = "Backend"
	backendStable.version = "1.0"
	backendStable.x = 300
	backendStable.y = 200
	scene.backendStable = backendStable

	//Preview backend
	backendPreview := microserviceGFX{}
	backendPreview.color = "green"
	backendPreview.name = "Backend"
	backendPreview.version = "2.0"
	backendPreview.x = 300
	backendPreview.y = 600
	scene.backendPreview = backendPreview

	//Production Queue
	queueStable := queueGFX{}
	queueStable.color = "red"
	queueStable.name = "production"
	queueStable.description = "RabbitMQ"
	//Just some starting values on the right of the diagram
	queueStable.x = 900
	queueStable.y = 100
	scene.queueStable = queueStable

	//Preview Queue
	queuePreview := queueGFX{}
	queuePreview.color = "orange"
	queuePreview.name = "preview"
	queuePreview.description = "RabbitMQ"
	//Mirrored downwards from stable version
	queuePreview.x = 900
	queuePreview.y = 500
	scene.queuePreview = queuePreview
}

func (scene *sceneGFX) renderQueue(queue queueGFX) {
	x := queue.x
	y := queue.y

	fontSize := 26
	svgTextOptions := fmt.Sprintf("font-size:%dpx;fill:black", fontSize)
	svgDiskOptions := fmt.Sprintf("fill:%s;stroke:black", queue.color)

	startingY := y + (diskSpacing * numberOfDisks)
	for i := 0; i < numberOfDisks; i++ {
		scene.canvas.Ellipse(x, startingY, queueDiskWidth, queueDiskHeight, svgDiskOptions, `id="frontendPreview"`)
		startingY = startingY - diskSpacing
	}

	//Description on top of the disks
	scene.canvas.Text(x-queueDiskWidth, y-margin, queue.description, svgTextOptions)

	//queue name is to the right of the disks
	scene.canvas.Text(x+queueDiskWidth+margin, y+(diskSpacing*numberOfDisks/2), queue.name, svgTextOptions)

}

func (scene *sceneGFX) renderMicroservice(ms microserviceGFX) {
	x := ms.x
	y := ms.y

	svgColorOptions := fmt.Sprintf("fill:%s;stroke:black", ms.color)
	svgIdOptions := fmt.Sprintf("id='%s'", ms.name)

	scene.canvas.Rect(x, y, rolloutWidth, rolloutHeight, svgColorOptions, svgIdOptions)

	//Description on top of the service
	fontSize := 26
	svgTextOptions := fmt.Sprintf("font-size:%dpx;fill:black", fontSize)
	scene.canvas.Text(x, y-margin, ms.name, svgTextOptions)

	//Version number inside the box
	scene.canvas.Text(x+(rolloutWidth/2), y+(rolloutHeight/2), ms.version, "text-anchor:middle;font-size:40px;fill:white")

}

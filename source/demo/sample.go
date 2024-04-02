package main

import (
	"log"
	"github.com/ajstarks/svgo"
	"net/http"
)

func main() {
	http.Handle("/circle", http.HandlerFunc(circle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "image/svg+xml")
  s := svg.New(w)
  s.Start(500, 500, `onload="initializeDraggableElements();"`, `xmlns:drag="http://www.codedread.com/dragsvg"`)
  s.Script("application/javascript", "http://www.codedread.com/dragsvg.js")
  s.Circle(250, 250, 125, `fill:none;stroke:black`,`drag:enable="true"`,`id="circle"`)
  s.Animate("#circle", "opacity", 0, 1, 3, 15)
  s.End()
}

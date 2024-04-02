package main

import (
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func (loanApp *LoanApplication) renderLiveDiagram(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(1200, 800)
	s.Circle(250, 250, 125, `fill:none;stroke:black`, `id="circle"`)
	s.Animate("#circle", "opacity", 0, 1, 3, 15)
	s.End()
}

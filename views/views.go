package views

import (
	_ "embed"

	"github.com/aymerick/raymond"
)

//go:embed index.html
var htmlIndex string

var indexTemplate = raymond.MustParse(htmlIndex)

type IndexPage struct {
	Day          string
	Going        bool
	Returning    bool
	EndOfShift   int
	EndOfDay     int
	StartOfDay   int
	DisableGoing bool
}

func RenderIndex(page IndexPage) string {
	return indexTemplate.MustExec(page)
}

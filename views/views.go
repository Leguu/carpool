package views

import (
	_ "embed"

	"github.com/aymerick/raymond"
)

//go:embed index.html
var htmlIndex string

var IndexTemplate = raymond.MustParse(htmlIndex)

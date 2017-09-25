package main

import (
	"gopkg.in/bblfsh/sdk.v1/protocol/driver"

	"github.com/bblfsh/ocaml-driver/driver/normalizer"
)

var version string
var build string

func main() {
	d := driver.Driver{
		Version:       version,
		Build:         build,
		ParserBuilder: normalizer.ParserBuilder,
		Annotate:      normalizer.AnnotationRules,
	}
	d.Exec()
}

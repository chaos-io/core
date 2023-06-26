package swaggerui_test

import (
	"net/http"

	swaggerui2 "github.com/chaos-io/core/httputil/swaggerui"
)

func Example_withJsonScheme() {
	swaggerScheme := []byte(`
{
  "openapi": "3.0.0",
  "info": {
    "title": "Sample API",
    "version": "0.1.9",
    "description": "Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML"
  }
}
`)

	http.Handle("/", http.FileServer(
		swaggerui2.NewFileSystem(swaggerui2.WithJSONScheme(swaggerScheme)),
	))
}

func Example_withYamlScheme() {
	swaggerScheme := []byte(`
---
openapi: 3.0.0
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9
`)

	http.Handle("/", http.FileServer(
		swaggerui2.NewFileSystem(swaggerui2.WithYAMLScheme(swaggerScheme)),
	))
}

func Example_withRemoteScheme() {
	http.Handle("/", http.FileServer(
		swaggerui2.NewFileSystem(
			swaggerui2.WithRemoteScheme("/my/scheme/path.yaml"),
		),
	))
}

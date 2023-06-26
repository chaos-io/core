package resource_test

import (
	"net/http"

	"github.com/chaos-io/core/httputil/resource"
)

func Example_stdlib() {
	uriPath := "/static/"
	http.Handle(uriPath, http.StripPrefix(uriPath, http.FileServer(resource.Dir("/static/"))))
}

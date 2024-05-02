package ctx

import (
	"net/http"
)

type requestContext struct {
	w       http.ResponseWriter
	r       *http.Request
	status  uint
	headers http.Header
}

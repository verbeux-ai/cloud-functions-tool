package cloud_functions_tool

import (
	"context"
	"mime/multipart"
	"net/http"
)

type Context interface {
	Context() context.Context
	JSON(data interface{}) Context
	Headers(header http.Header) Context
	Status(int) Context
	GetFormFile(string) (*multipart.FileHeader, error)
	Unmarshall(target interface{}) error
}

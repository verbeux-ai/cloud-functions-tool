package cloud_functions_tool

import (
	"mime/multipart"
	"net/http"
)

type Context interface {
	JSON(data interface{})
	Headers(header http.Header) Context
	Status(int) Context
	GetFormFile(string) (*multipart.FileHeader, error)
	Unmarshall(target interface{}) error
}

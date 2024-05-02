package ctx

import (
	"mime/multipart"
	"net/http"
)

type Context interface {
	JSON(data interface{}) error
	Headers(header http.Header) Context
	Status(uint) Context
	GetFormFile(string) (*multipart.FileHeader, error)
}

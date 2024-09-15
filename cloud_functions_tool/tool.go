package cloud_functions_tool

import (
	"encoding/json"
	"mime/multipart"
	"net/http"

	"go.uber.org/zap"
)

func NewContext(w http.ResponseWriter, r *http.Request) Context {
	return &requestContext{w, r, 0, make(http.Header)}
}

func (s *requestContext) JSON(data interface{}) {
	s.w.Header().Set("Content-Type", "application/json")
	if s.status == 0 {
		s.w.WriteHeader(http.StatusOK)
	}

	err := json.NewEncoder(s.w).Encode(data)
	if err != nil {
		zap.L().Error("failed to encode json to writer", zap.Error(err))
		if _, err = s.w.Write([]byte(`{"message": "fatal", "error": "fatal"}`)); err != nil {
			zap.L().Panic("failed to return response", zap.Error(err))
		}
	}
}

func (s *requestContext) Unmarshall(target interface{}) error {
	return json.NewDecoder(s.r.Body).Decode(&target)
}

func (s *requestContext) Headers(m http.Header) Context {
	s.headers = m
	for k, v := range m {
		for _, hV := range v {
			s.w.Header().Add(k, hV)
		}
	}
	return s
}

func (s *requestContext) Status(i int) Context {
	if i > 0 {
		s.w.WriteHeader(i)
	}
	s.status = i
	return s
}

func (s *requestContext) GetFormFile(key string) (*multipart.FileHeader, error) {
	err := s.r.ParseMultipartForm(100 << 20)
	if err != nil {
		return nil, err
	}

	file, fileHeader, err := s.r.FormFile(key)
	if err != nil {
		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			zap.L().Error("failed to close file", zap.Error(err))
		}
	}(file)

	return fileHeader, nil
}

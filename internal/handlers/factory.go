package handlers

import (
	"go-employee-web-server/internal/api"
	"go-employee-web-server/internal/data"
	"net/http"
)

// Factory is an interface for creating handlers
type Factory interface {
	MakeEmployeesHandler() http.HandlerFunc
	MakeViewHandler() http.HandlerFunc
	MakeEditHandler() http.HandlerFunc
	MakeAddHandler() http.HandlerFunc
}

// HandlerFactory is a factory for creating HTTP handlers
type HandlerFactory struct {
	Storage   data.Storage
	APIClient api.APIClient
}

// NewHandlerFactory creates a new instance of HandlerFactory
func NewHandlerFactory(storage data.Storage, apiClient api.APIClient) *HandlerFactory {
	return &HandlerFactory{
		Storage:   storage,
		APIClient: apiClient,
	}
}

// MakeEmployeesHandler creates a new employees handler
func (f *HandlerFactory) MakeEmployeesHandler() http.HandlerFunc {
	return EmployeesHandler(f.Storage, f.APIClient)
}

// MakeViewHandler creates a new view handler
func (f *HandlerFactory) MakeViewHandler() http.HandlerFunc {
	return ViewHandler(f.Storage)
}

// MakeEditHandler creates a new edit handler
func (f *HandlerFactory) MakeEditHandler() http.HandlerFunc {
	return EditHandler(f.Storage)
}

func (f *HandlerFactory) MakeAddHandler() http.HandlerFunc {
	return AddHandler(f.Storage)
}

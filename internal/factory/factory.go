package factory

import (
	"github.com/go-redis/redis/v8"
	"go-employee-web-server/internal/api"
	"go-employee-web-server/internal/data"
	"go-employee-web-server/internal/handlers"
	"net/http"
)

// Factory is an interface for creating handlers
type Factory interface {
	MakeEmployeesHandler() http.HandlerFunc
	MakeViewHandler() http.HandlerFunc
	MakeEditHandler() http.HandlerFunc
	MakeAddHandler() http.HandlerFunc
	MakeLoginHandler() http.HandlerFunc
	MakePingHandler() http.HandlerFunc
	MakeTopHandler() http.HandlerFunc
	MakeCountHandler() http.HandlerFunc
}

// HandlerFactory is a factory for creating HTTP handlers
type HandlerFactory struct {
	Storage     data.Storage
	APIClient   api.APIClient
	RedisClient *redis.Client
}

// NewHandlerFactory creates a new instance of HandlerFactory
func NewHandlerFactory(storage data.Storage, apiClient api.APIClient, redisClient *redis.Client) *HandlerFactory {
	return &HandlerFactory{
		Storage:     storage,
		APIClient:   apiClient,
		RedisClient: redisClient,
	}
}

// MakeEmployeesHandler creates a new employees handler
func (f *HandlerFactory) MakeEmployeesHandler() http.HandlerFunc {
	return handlers.EmployeesHandler(f.Storage, f.APIClient)
}

// MakeViewHandler creates a new view handler
func (f *HandlerFactory) MakeViewHandler() http.HandlerFunc {
	return handlers.ViewHandler(f.Storage)
}

// MakeEditHandler creates a new edit handler
func (f *HandlerFactory) MakeEditHandler() http.HandlerFunc {
	return handlers.EditHandler(f.Storage)
}

func (f *HandlerFactory) MakeAddHandler() http.HandlerFunc {
	return handlers.AddHandler(f.Storage)
}

func (f *HandlerFactory) MakeLoginHandler() http.HandlerFunc {
	return handlers.LoginHandler(f.RedisClient)
}

func (f *HandlerFactory) MakePingHandler() http.HandlerFunc {
	return handlers.PingHandler(f.RedisClient)
}

func (f *HandlerFactory) MakeTopHandler() http.HandlerFunc {
	return handlers.TopHandler(f.RedisClient)
}

func (f *HandlerFactory) MakeCountHandler() http.HandlerFunc {
	return handlers.CountHandler(f.RedisClient)
}

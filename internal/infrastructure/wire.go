package infrastructure

import (
	. "eventure_backend/internal/infrastructure/configuration"
	"github.com/google/wire"
)

// InitializeApp initializes all dependencies for the application
func InitializeApp() *Server {
	wire.Build(
		ProvideGormDB,
		ProvideElasticClient,
	)
	return &Server{}
}

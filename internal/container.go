package internal

import (
	. "eventure_backend/internal/api"
	. "eventure_backend/internal/infrastructure/database"
	"github.com/google/wire"
)

var ContainerSet = wire.NewSet(InitializeDatabase, NewRouter)

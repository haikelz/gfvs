package server

import (
	"gfvs/pkg/server/routes"
)

func (s *FiberApp) RegisterFiberRoutes() {
	routes.HomeRoute(s)
}

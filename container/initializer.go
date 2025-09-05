package container

import (
	addon_controller_v1 "confluence-checkout/internal/feature-add-on/controller/v1"
	addon_service "confluence-checkout/internal/feature-add-on/service"
	"confluence-checkout/internal/infrastructure/config"
	"confluence-checkout/internal/infrastructure/persistence"
)

type Initializer struct {
	AddOnService      *addon_service.AddOnServiceHandler
	AddOnControllerV1 *addon_controller_v1.AddOnController
}

func NewInitializer(env config.Env) *Initializer {
	initializer := &Initializer{}

	postgresInfra := persistence.NewPostgresInfraHandler()

	initializer.AddOnService = addon_service.NewAddOnServiceHandler(postgresInfra)
	initializer.AddOnControllerV1 = addon_controller_v1.NewAddOnController(initializer.AddOnService)

	return initializer
}

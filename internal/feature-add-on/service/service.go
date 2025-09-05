package service

import (
	dto_v1 "confluence-checkout/internal/feature-add-on/dto/v1"
	"confluence-checkout/internal/infrastructure/persistence"
	"confluence-checkout/pkg/pkg_dto"
	"context"
)

type AddOnService interface {
	Create(ctx context.Context, data dto_v1.AddOnRequest) (res pkg_dto.IdResponse, traceID string, statusCode int, err error)
}

type AddOnServiceHandler struct {
	Database persistence.DatabaseInfra
}

func NewAddOnServiceHandler(database persistence.DatabaseInfra) *AddOnServiceHandler {
	return &AddOnServiceHandler{
		Database: database,
	}
}

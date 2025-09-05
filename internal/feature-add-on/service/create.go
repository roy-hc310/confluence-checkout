package service

import (
	"confluence-checkout/internal/feature-add-on/domain"
	dto_v1 "confluence-checkout/internal/feature-add-on/dto/v1"
	"confluence-checkout/pkg/pkg_dto"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (h *AddOnServiceHandler) Create(ctx context.Context, data dto_v1.AddOnRequest) (res pkg_dto.IdResponse, traceID string, statusCode int, err error) {

	addOn := domain.AddOn{}

	err = copier.Copy(&addOn, &data)
	if err != nil {
		return res, traceID, http.StatusOK, err
	}

	addOn.UUID, err = uuid.NewV7()
	if err != nil {
		return res, traceID, http.StatusOK, err
	}

	fmt.Println("howhefoh") // param is a struct, then loop through each ele to add to interface, then prepare insert query
	addOnString := fmt.Sprintf(
		`SELECT name, age FROM %s WHERE name=$1;`, "\"checkout\".\"user\"",
	)

	name := "testname"

	rows, err := h.Database.Query(ctx, addOnString, []interface{}{name})
	if err != nil {
		return res, traceID, 500, err
	}
	defer rows.Close()
	println(rows)

	return res, traceID, statusCode, nil
}

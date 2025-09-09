package service

import (
	"confluence-checkout/internal/feature-add-on/domain"
	dto_v1 "confluence-checkout/internal/feature-add-on/dto/v1"
	"confluence-checkout/pkg/pkg_dto"
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

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

	result := tempp(&addOn)
	fmt.Println(result)
	// fmt.Println(result["products"])

	// fmt.Println("howhefoh") // param is a struct, then loop through each ele to add to interface, then prepare insert query
	// addOnString := fmt.Sprintf(
	// 	`SELECT name, age FROM %s WHERE name=$1;`, "\"checkout\".\"user\"",
	// )

	// name := "testname"

	// rows, err := h.Database.Query(ctx, addOnString, []interface{}{name})
	// if err != nil {
	// 	return res, traceID, 500, err
	// }
	// defer rows.Close()
	// println(rows)

	return res, traceID, statusCode, nil
}

// func temping(ctx context.Context, tx pgx.Tx, data interface{}, tableName string) map[string]interface{} {}

func temp2(data map[string]interface{}) {

}

func tempp(data interface{}) map[string]interface{} {

	result := make(map[string]interface{})
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if !value.CanInterface() {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		tagName := strings.Split(tag, ",")[0]
		if tagName == "" {
			continue
		}

		switch value.Kind() {
		case reflect.Pointer:
			if value.IsNil() {
				result[tagName] = nil
				continue
			}
			result[tagName] = value.Elem().Interface()

		case reflect.Struct:
			result[tagName] = tempp(value.Interface())
		case reflect.Slice:
			var sliceVals []interface{}
			for j := 0; j < value.Len(); j++ {
				elem := value.Index(j)
				if elem.Kind() == reflect.Pointer && !elem.IsNil() {
					elem = elem.Elem()
				}
				hehe := elem.Interface()
				fmt.Println(hehe)
				sliceVals = append(sliceVals, tempp(elem.Interface()))
			}
			result[tagName] = sliceVals
		default:
			result[tagName] = value.Interface()
		}
	}

	return result
}

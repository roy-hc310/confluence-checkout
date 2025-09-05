package domain

import (
	"confluence-checkout/pkg/pkg_entity"
	"time"
)

type AddOn struct {
	pkg_entity.BaseEntity
	Name                      string    `json:"name" validate:"required"`
	Code                      string    `json:"code" validate:"required"`
	ShopID                    string    `json:"shop_id" validate:"required"`
	Active                    bool      `json:"active"`
	StartTime                 time.Time `json:"start_time" validate:"required"`
	EndTime                   time.Time `json:"end_time" validate:"required"`
	UsageQuantity             *int      `json:"usage_quantity"`
	UsageQuantityRemaining    *int      `json:"usage_quantity_remaining"`
	UsageLimitPerUser         *int      `json:"usage_limit_per_user"`
	ConditionalMinSpendAmount float64   `json:"conditional_min_spend_amount" validate:"required,gte=0"`
	Products                  []Product `json:"products" validate:"required,min=1,dive"`
}

type Product struct {
	pkg_entity.BaseEntity
	AddOnID  string `json:"add_on_id"`
	SKU      string `json:"sku" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gte=1"`
}

type AddOnUsageLimitPerUser struct {
	pkg_entity.BaseEntity
	AddOnID    string `json:"add_on_id"`
	CustomerID string `json:"customer_id"`
	Quantity   int    `json:"quantity"`
}

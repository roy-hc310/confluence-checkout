package container

import (
	add_on "confluence-checkout/internal/feature-add-on"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine, initializer *Initializer) {

	route := g.Group("/api")

	addOnRouter := route.Group("/add-on")
	add_on.AddOnRoute(addOnRouter, initializer.AddOnControllerV1)
}

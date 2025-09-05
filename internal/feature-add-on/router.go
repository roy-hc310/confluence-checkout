package add_on

import (
	addon_controller_v1 "confluence-checkout/internal/feature-add-on/controller/v1"

	"github.com/gin-gonic/gin"
)

func AddOnRoute(addOnRoute *gin.RouterGroup, controller *addon_controller_v1.AddOnController) {
	v1Route := addOnRoute.Group("/v1")
	{
		v1Route.POST("/create", controller.Create)
	}

}

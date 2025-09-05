package v1

import (
	dto "confluence-checkout/internal/feature-add-on/dto/v1"
	"confluence-checkout/internal/feature-add-on/service"
	"confluence-checkout/pkg/pkg_dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddOnController struct {
	AddOnService service.AddOnService
}

func NewAddOnController(addOnServiceHandler service.AddOnService) *AddOnController {
	return &AddOnController{
		AddOnService: addOnServiceHandler,
	}
}

func (c *AddOnController) Create(g *gin.Context) {
	// Implement the logic to handle the create request
	res := pkg_dto.BaseResponse{}

	body := dto.AddOnRequest{}
	err := g.ShouldBindJSON(&body)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, traceID, statusCode, err := c.AddOnService.Create(g.Request.Context(), body)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		g.AbortWithStatusJSON(statusCode, res)
		return
	}

	res.Data = data
	res.TraceID = traceID
	res.Succeeded = true
	g.JSON(statusCode, res)
}

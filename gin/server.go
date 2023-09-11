package gin

import (
	"ondc/api"
	"ondc/model"
	"ondc/model/operations"

	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine, handler api.SellerApp) {

	engine.POST("/search", func(c *gin.Context) {
		request := operations.SearchRequest{}
		response, err := handler.Search(&request)

		if err != nil {
			c.JSON(err.(model.ApiError).HttpStatus, err)
		} else {
			c.JSON(200, response)
		}
	})
}

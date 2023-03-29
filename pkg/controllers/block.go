package controllers

import (
	http_response "web3-with-go/pkg/http/response"
	"web3-with-go/pkg/services"

	"github.com/gin-gonic/gin"
)

type IBlockRoutesController interface {
	GetCurrentBlock(c *gin.Context)
}

type BlockRoutesController struct{}

var BlockController IBlockRoutesController = &BlockRoutesController{}

// GetCurrentBlock godoc
// @Summary Get current block of a blackchain
// @Tags block
// @Accept  json
// @Produce  json
// @Param protocol query string true " "
// @Success 200
// @Router /current-block [get]
func (b *BlockRoutesController) GetCurrentBlock(c *gin.Context) {
	protocol := c.Query("protocol")
	if protocol == "" {
		http_response.SendBadRequest(c, "invalid protocol query string")
		return
	}
	block, err := services.GetCurrentBlock(protocol)
	if err != nil {
		http_response.SendJsonHttpError(c, err)
		return
	}

	http_response.SendJson(c, 200, block)
}

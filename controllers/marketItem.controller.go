package controllers

import (
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type MarketItemController struct {
	MarketItemService services.MarketItemService
}

func NewMarketItemController(marketItemservice services.MarketItemService) MarketItemController {
	return MarketItemController{
		MarketItemService: marketItemservice,
	}
}

func (mc *MarketItemController) CreateMarketItem(ctx *gin.Context) {
	var marketItem models.MarketItem
	if err := ctx.ShouldBindJSON(&marketItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := mc.MarketItemService.Create(&marketItem)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (mc *MarketItemController) GetMarketItem(ctx *gin.Context) {
	var marketItemId string = ctx.Param("id")
	market, err := mc.MarketItemService.Show(&marketItemId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, market)
}

func (mc *MarketItemController) List(ctx *gin.Context) {
	marketItems, err := mc.MarketItemService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, marketItems)
}

func (mc *MarketItemController) UpdateMarketItem(ctx *gin.Context) {
	var marketItem models.MarketItem
	if err := ctx.ShouldBindJSON(&marketItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	marketItemUpdate, err := mc.MarketItemService.Update(&marketItem)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    marketItemUpdate,
	})
}

func (mc *MarketItemController) DeleteMarketItem(ctx *gin.Context) {
	var marketItem_id string = ctx.Param("id")
	err := mc.MarketItemService.Delete(&marketItem_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (mc *MarketItemController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/marketItems")
	userroute.POST("", mc.CreateMarketItem)
	userroute.GET("/:id", mc.GetMarketItem)
	userroute.GET("", mc.List)
	userroute.PATCH("", mc.UpdateMarketItem)
	userroute.DELETE("/:id", mc.DeleteMarketItem)
}

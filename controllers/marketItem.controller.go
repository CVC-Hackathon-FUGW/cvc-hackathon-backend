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
	_, err := mc.MarketItemService.Update(&marketItem)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
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

func (mc *MarketItemController) FindByAddress(ctx *gin.Context) {
	var address string = ctx.Param("address")
	markets, err := mc.MarketItemService.FindByAddress(&address)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, markets)
}

func (mc *MarketItemController) BuyMarketItem(ctx *gin.Context) {
	var marketItem_id string = ctx.Param("id")
	err := mc.MarketItemService.BuyMarketItem(&marketItem_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (mc *MarketItemController) OfferMarketItem(ctx *gin.Context) {
	var marketItem_id string = ctx.Param("id")
	err := mc.MarketItemService.OfferMarketItem(&marketItem_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (mc *MarketItemController) RegisterRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/marketItems")
	route.POST("", mc.CreateMarketItem)
	route.GET("/:id", mc.GetMarketItem)
	route.GET("", mc.List)
	route.PATCH("", mc.UpdateMarketItem)
	route.DELETE("/:id", mc.DeleteMarketItem)
	route.GET("/address/:address", mc.FindByAddress)
	route.POST("/buy/:id", mc.BuyMarketItem)
	route.POST("/offer/:id", mc.OfferMarketItem)
}

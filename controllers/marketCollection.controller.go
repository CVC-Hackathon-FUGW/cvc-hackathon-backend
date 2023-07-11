package controllers

import (
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type MarketCollectionController struct {
	MarketCollectionService services.MarketCollectionService
}

func NewMarketCollectionController(marketCollectionservice services.MarketCollectionService) MarketCollectionController {
	return MarketCollectionController{
		MarketCollectionService: marketCollectionservice,
	}
}

func (mc *MarketCollectionController) CreateMarketCollection(ctx *gin.Context) {
	var marketCollection models.MarketCollection
	if err := ctx.ShouldBindJSON(&marketCollection); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := mc.MarketCollectionService.Create(&marketCollection)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (mc *MarketCollectionController) GetMarketCollection(ctx *gin.Context) {
	var marketCollectionId string = ctx.Param("id")
	market, err := mc.MarketCollectionService.Show(&marketCollectionId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, market)
}

func (mc *MarketCollectionController) List(ctx *gin.Context) {
	params := enum.MarketCollectionsParams{
		Name: ctx.Query("name"),
	}
	marketCollections, err := mc.MarketCollectionService.List(params)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, marketCollections)
}

func (mc *MarketCollectionController) UpdateMarketCollection(ctx *gin.Context) {
	var marketCollection models.MarketCollection
	if err := ctx.ShouldBindJSON(&marketCollection); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	marketCollectionUpdate, err := mc.MarketCollectionService.Update(&marketCollection)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    marketCollectionUpdate,
	})
}

func (mc *MarketCollectionController) DeleteMarketCollection(ctx *gin.Context) {
	var MarketCollection_id string = ctx.Param("id")
	err := mc.MarketCollectionService.Delete(&MarketCollection_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (mc *MarketCollectionController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/marketCollections")
	userroute.POST("", mc.CreateMarketCollection)
	userroute.GET("/:id", mc.GetMarketCollection)
	userroute.GET("", mc.List)
	userroute.PATCH("", mc.UpdateMarketCollection)
	userroute.DELETE("/:id", mc.DeleteMarketCollection)
}

package controllers

import (
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type SellerController struct {
	SellerService services.SellerService
}

func NewSellerController(Sellerservice services.SellerService) SellerController {
	return SellerController{
		SellerService: Sellerservice,
	}
}

func (uc *SellerController) CreateSeller(ctx *gin.Context) {
	var seller models.Seller
	if err := ctx.ShouldBindJSON(&seller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.SellerService.Create(&seller)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *SellerController) GetSellerByAddress(ctx *gin.Context) {
	var addr string = ctx.Param("address")
	seller, err := uc.SellerService.Show(&addr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, seller)
}

func (uc *SellerController) List(ctx *gin.Context) {
	sellers, err := uc.SellerService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sellers)
}

func (uc *SellerController) RegisterRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/sellers")
	route.POST("", uc.CreateSeller)
	route.GET("/address/:address", uc.GetSellerByAddress)
	route.GET("", uc.List)
}

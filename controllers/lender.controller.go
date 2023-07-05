package controllers

import (
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type LenderController struct {
	LenderService services.LenderService
}

func NewLenderController(lenderservice services.LenderService) LenderController {
	return LenderController{
		LenderService: lenderservice,
	}
}

func (uc *LenderController) CreateLender(ctx *gin.Context) {
	var lender models.Lender
	if err := ctx.ShouldBindJSON(&lender); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.LenderService.Create(&lender)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *LenderController) GetLender(ctx *gin.Context) {
	var lenderID string = ctx.Param("id")
	user, err := uc.LenderService.Show(&lenderID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *LenderController) List(ctx *gin.Context) {
	lenders, err := uc.LenderService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, lenders)
}

func (uc *LenderController) UpdateLender(ctx *gin.Context) {
	var lender models.Lender
	if err := ctx.ShouldBindJSON(&lender); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	lenderupdate, err := uc.LenderService.Update(&lender)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    lenderupdate,
	})
}

func (uc *LenderController) DeleteLender(ctx *gin.Context) {
	var lender_id string = ctx.Param("id")
	err := uc.LenderService.Delete(&lender_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *LenderController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/lenders")
	userroute.POST("", uc.CreateLender)
	userroute.GET("/:id", uc.GetLender)
	userroute.GET("", uc.List)
	userroute.PATCH("", uc.UpdateLender)
	userroute.DELETE("/:id", uc.DeleteLender)
}

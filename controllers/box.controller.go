package controllers

import (
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BoxController struct {
	BoxService services.BoxService
}

func NewBoxController(boxservice services.BoxService) BoxController {
	return BoxController{
		BoxService: boxservice,
	}
}

func (uc *BoxController) CreateBox(ctx *gin.Context) {
	var box models.Box
	if err := ctx.ShouldBindJSON(&box); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.BoxService.Create(&box)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BoxController) GetBox(ctx *gin.Context) {
	var boxId string = ctx.Param("id")
	box, err := uc.BoxService.Show(&boxId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, box)
}

func (uc *BoxController) List(ctx *gin.Context) {
	boxes, err := uc.BoxService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, boxes)
}

func (uc *BoxController) UpdateBox(ctx *gin.Context) {
	var box models.Box
	if err := ctx.ShouldBindJSON(&box); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err := uc.BoxService.Update(&box)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BoxController) DeleteBox(ctx *gin.Context) {
	var boxId string = ctx.Param("id")
	err := uc.BoxService.Delete(&boxId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BoxController) FindByAddress(ctx *gin.Context) {
	var address string = ctx.Param("address")
	boxes, err := uc.BoxService.FindByAddress(&address)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, boxes)
}

func (uc *BoxController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/box", uc.CreateBox)
	router.GET("/box/:id", uc.GetBox)
	router.GET("/box", uc.List)
	router.PATCH("/box", uc.UpdateBox)
	router.DELETE("/box/:id", uc.DeleteBox)
	router.GET("/box/address/:address", uc.FindByAddress)
}

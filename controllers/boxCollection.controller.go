package controllers

import (
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BoxCollectionController struct {
	BoxCollectionService services.BoxCollectionService
}

func NewBoxCollectionController(boxCollectionService services.BoxCollectionService) BoxCollectionController {
	return BoxCollectionController{
		BoxCollectionService: boxCollectionService,
	}
}

func (uc *BoxCollectionController) CreateBoxCollection(ctx *gin.Context) {
	var boxCollection models.BoxCollection
	if err := ctx.ShouldBindJSON(&boxCollection); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.BoxCollectionService.Create(&boxCollection)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BoxCollectionController) GetBoxCollection(ctx *gin.Context) {
	var boxCollectionId string = ctx.Param("id")
	boxCollection, err := uc.BoxCollectionService.Show(&boxCollectionId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, boxCollection)
}

func (uc *BoxCollectionController) List(ctx *gin.Context) {
	boxCollections, err := uc.BoxCollectionService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, boxCollections)
}

func (uc *BoxCollectionController) UpdateBoxCollection(ctx *gin.Context) {
	var boxCollection models.BoxCollection
	if err := ctx.ShouldBindJSON(&boxCollection); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := uc.BoxCollectionService.Update(&boxCollection)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BoxCollectionController) DeleteBoxCollection(ctx *gin.Context) {
	var boxCollectionId string = ctx.Param("id")
	err := uc.BoxCollectionService.Delete(&boxCollectionId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BoxCollectionController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/boxCollection", uc.CreateBoxCollection)
	router.GET("/boxCollection/:id", uc.GetBoxCollection)
	router.GET("/boxCollection", uc.List)
	router.PATCH("/boxCollection", uc.UpdateBoxCollection)
	router.DELETE("/boxCollection/:id", uc.DeleteBoxCollection)
}

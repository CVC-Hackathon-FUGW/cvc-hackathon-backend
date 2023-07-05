package controllers

import (
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type BorrowerController struct {
	borrowerService services.BorrowerService
}

func NewBorrowerController(borrowerService services.BorrowerService) BorrowerController {
	return BorrowerController{
		borrowerService: borrowerService,
	}
}

func (uc *BorrowerController) CreateBorrower(ctx *gin.Context) {
	var borrower models.Borrower
	if err := ctx.ShouldBindJSON(&borrower); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.borrowerService.Create(&borrower)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BorrowerController) GetBorrower(ctx *gin.Context) {
	var borrowerID string = ctx.Param("id")
	user, err := uc.borrowerService.Show(&borrowerID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *BorrowerController) List(ctx *gin.Context) {
	borrowers, err := uc.borrowerService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, borrowers)
}

func (uc *BorrowerController) UpdateBorrower(ctx *gin.Context) {
	var borrower models.Borrower
	if err := ctx.ShouldBindJSON(&borrower); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	borrowerupdate, err := uc.borrowerService.Update(&borrower)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    borrowerupdate,
	})
}

func (uc *BorrowerController) DeleteBorrower(ctx *gin.Context) {
	var borrower_id string = ctx.Param("id")
	err := uc.borrowerService.Delete(&borrower_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *BorrowerController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/borrowers")
	userroute.POST("", uc.CreateBorrower)
	userroute.GET("/:id", uc.GetBorrower)
	userroute.GET("", uc.List)
	userroute.PATCH("", uc.UpdateBorrower)
	userroute.DELETE("/:id", uc.DeleteBorrower)
}

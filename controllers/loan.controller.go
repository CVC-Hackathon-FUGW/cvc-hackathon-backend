package controllers

import (
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type LoanController struct {
	LoanService services.LoanService
}

func NewLoanController(loanservice services.LoanService) LoanController {
	return LoanController{
		LoanService: loanservice,
	}
}

func (uc *LoanController) CreateLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.LoanService.Create(&loan)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *LoanController) GetLoan(ctx *gin.Context) {
	var LoanId string = ctx.Param("id")
	user, err := uc.LoanService.Show(&LoanId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *LoanController) List(ctx *gin.Context) {
	Loans, err := uc.LoanService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Loans)
}

func (uc *LoanController) UpdateLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	loanupdate, err := uc.LoanService.Update(&loan)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    loanupdate,
	})
}

func (uc *LoanController) DeleteLoan(ctx *gin.Context) {
	var loan_id string = ctx.Param("id")
	err := uc.LoanService.Delete(&loan_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *LoanController) MaxAMount(ctx *gin.Context) {
	var pool_id string = ctx.Param("id")
	loan, err := uc.LoanService.MaxAmount(&pool_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    loan,
	})
}

func (uc *LoanController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/loans")
	userroute.POST("", uc.CreateLoan)
	userroute.GET("/:id", uc.GetLoan)
	userroute.GET("", uc.List)
	userroute.PATCH("", uc.UpdateLoan)
	userroute.DELETE("/:id", uc.DeleteLoan)
	userroute.GET("/pool/:id/max-amount", uc.MaxAMount)

}

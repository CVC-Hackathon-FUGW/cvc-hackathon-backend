package controllers

import (
	"fmt"
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
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
	_, err := uc.LoanService.Update(&loan)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (uc *LoanController) DeleteLoan(ctx *gin.Context) {
	params := enum.LoanParams{
		WithPool: ctx.GetBool("with-pool"),
	}

	var loan_id string = ctx.Param("id")
	if params.WithPool {
		err := uc.LoanService.DeleteWithUpdatePool(&loan_id)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		return
	}

	err := uc.LoanService.Delete(&loan_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *LoanController) MaxAMount(ctx *gin.Context) {
	var poolID string = ctx.Param("id")
	loan, err := uc.LoanService.MaxAmount(&poolID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    loan[0].Amount,
	})
}

func (uc *LoanController) CountLoan(ctx *gin.Context) {
	var pool_id string = ctx.Param("id")
	loan, err := uc.LoanService.CountLoans(&pool_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    fmt.Sprintf("%d of %d offers taken", loan.TotalLoanGot, loan.TotalLoanInPool),
	})
}

func (uc *LoanController) BorrowserTakeLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.LoanService.BorrowserTakeLoan(&loan)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
func (uc *LoanController) RegisterRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/loans")
	route.POST("", uc.CreateLoan)
	route.GET("/:id", uc.GetLoan)
	route.GET("", uc.List)
	route.PATCH("", uc.UpdateLoan)
	route.DELETE("/:id", uc.DeleteLoan)
	route.GET("/pool/:id/max-amount", uc.MaxAMount)
	route.GET("/pool/:id/count", uc.CountLoan)
	route.PATCH("/borrower-take-loan", uc.BorrowserTakeLoan)
}

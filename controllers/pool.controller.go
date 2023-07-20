package controllers

import (
	"fmt"
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type PoolController struct {
	PoolService services.PoolService
}

func NewPool(poolservice services.PoolService) PoolController {
	return PoolController{
		PoolService: poolservice,
	}
}

func (uc *PoolController) CreatePool(ctx *gin.Context) {
	var pool models.Pool
	if err := ctx.ShouldBindJSON(&pool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.PoolService.Create(&pool)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *PoolController) GetPool(ctx *gin.Context) {
	var poolID string = ctx.Param("id")
	pool, err := uc.PoolService.Show(&poolID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	loanMaxAmount, err := uc.PoolService.MaxAmount(&poolID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	loanCount, err := uc.PoolService.CountLoans(&poolID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	poolwithloan := models.PoolWithLoanDetails{
		LoanCount:     fmt.Sprintf("%d of %d offers taken", loanCount.TotalLoanGot, loanCount.TotalLoanInPool),
		LoanMaxAmount: loanMaxAmount[0].Amount,
		Pool:          pool,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    poolwithloan,
	})
}

func (uc *PoolController) List(ctx *gin.Context) {
	params := enum.PoolParams{
		Name: ctx.Query("name"),
	}

	pools, err := uc.PoolService.List(params)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pools)
}

func (uc *PoolController) UpdatePool(ctx *gin.Context) {
	var pool models.Pool
	if err := ctx.ShouldBindJSON(&pool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	poolupdate, err := uc.PoolService.Update(&pool)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    poolupdate,
	})
}

func (uc *PoolController) DeletePool(ctx *gin.Context) {
	var pool_id string = ctx.Param("id")
	err := uc.PoolService.Delete(&pool_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *PoolController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/pools")
	userroute.POST("", uc.CreatePool)
	userroute.GET("/:id", uc.GetPool)
	userroute.GET("", uc.List)
	userroute.PATCH("", uc.UpdatePool)
	userroute.DELETE("/:id", uc.DeletePool)
}

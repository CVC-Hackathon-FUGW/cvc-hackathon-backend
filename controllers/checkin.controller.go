package controllers

import (
	"net/http"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
)

type CheckinController struct {
	CheckinService services.CheckinService
}

func NewCheckinController(Checkinservice services.CheckinService) CheckinController {
	return CheckinController{
		CheckinService: Checkinservice,
	}
}

func (uc *CheckinController) CreateCheckin(ctx *gin.Context) {
	var checkin models.Checkin
	if err := ctx.ShouldBindJSON(&checkin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.CheckinService.Create(&checkin)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *CheckinController) GetCheckin(ctx *gin.Context) {
	var checkinId string = ctx.Param("id")
	user, err := uc.CheckinService.Show(&checkinId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *CheckinController) List(ctx *gin.Context) {
	params := enum.CheckinParams{
		Name: ctx.Query("name"),
		Sort: ctx.Query("sort"),
	}

	Checkins, err := uc.CheckinService.List(params)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Checkins)
}

func (uc *CheckinController) UpdateCheckin(ctx *gin.Context) {
	var Checkin models.Checkin
	if err := ctx.ShouldBindJSON(&Checkin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	Checkinupdate, err := uc.CheckinService.Update(&Checkin)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    Checkinupdate,
	})
}

func (uc *CheckinController) DeleteCheckin(ctx *gin.Context) {
	var Checkin_id string = ctx.Param("id")
	err := uc.CheckinService.Delete(&Checkin_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *CheckinController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/checkins")
	userroute.POST("", uc.CreateCheckin)
	userroute.GET("/:id", uc.GetCheckin)
	userroute.GET("", uc.List)
	userroute.PATCH("", uc.UpdateCheckin)
	userroute.DELETE("/:id", uc.DeleteCheckin)
}

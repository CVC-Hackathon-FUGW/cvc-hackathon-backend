package controllers

import (
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ParticipantController struct {
	ParticipantService services.ParticipantService
}

func NewParticipantController(participantService services.ParticipantService) ParticipantController {
	return ParticipantController{
		ParticipantService: participantService,
	}
}

func (uc *ParticipantController) CreateParticipant(ctx *gin.Context) {
	var participant models.Participant
	if err := ctx.ShouldBindJSON(&participant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.ParticipantService.Create(&participant)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ParticipantController) GetParticipant(ctx *gin.Context) {
	var participantId string = ctx.Param("id")
	participant, err := uc.ParticipantService.Show(&participantId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, participant)
}

func (uc *ParticipantController) List(ctx *gin.Context) {
	participants, err := uc.ParticipantService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, participants)
}

func (uc *ParticipantController) UpdateParticipant(ctx *gin.Context) {
	var participant models.Participant
	if err := ctx.ShouldBindJSON(&participant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := uc.ParticipantService.Update(&participant)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ParticipantController) DeleteParticipant(ctx *gin.Context) {
	var participantId string = ctx.Param("id")
	err := uc.ParticipantService.Delete(&participantId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ParticipantController) Invest(ctx *gin.Context) {
	var participant models.Participant
	if err := ctx.ShouldBindJSON(&participant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.ParticipantService.Invest(&participant)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ParticipantController) FindByAddress(ctx *gin.Context) {
	var address string = ctx.Param("address")
	participants, err := uc.ParticipantService.FindByAddress(&address)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, participants)
}

func (uc *ParticipantController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/participant", uc.CreateParticipant)
	router.GET("/participant/:id", uc.GetParticipant)
	router.GET("/participant", uc.List)
	router.PATCH("/participant", uc.UpdateParticipant)
	router.DELETE("/participant/:id", uc.DeleteParticipant)
	router.POST("/participant/invest", uc.Invest)
	router.GET("/participant/address/:address", uc.FindByAddress)
}

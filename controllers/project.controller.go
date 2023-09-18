package controllers

import (
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectController struct {
	ProjectService services.ProjectService
}

func NewProjectController(projectservice services.ProjectService) ProjectController {
	return ProjectController{
		ProjectService: projectservice,
	}
}

func (uc *ProjectController) CreateProject(ctx *gin.Context) {
	var project models.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.ProjectService.Create(&project)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ProjectController) GetProject(ctx *gin.Context) {
	var projectId string = ctx.Param("id")
	project, err := uc.ProjectService.Show(&projectId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, project)
}

func (uc *ProjectController) List(ctx *gin.Context) {
	projects, err := uc.ProjectService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, projects)
}

func (uc *ProjectController) UpdateProject(ctx *gin.Context) {
	var project models.Project
	if err := ctx.ShouldBindJSON(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := uc.ProjectService.Update(&project)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ProjectController) DeleteProject(ctx *gin.Context) {
	var projectId string = ctx.Param("id")
	err := uc.ProjectService.Delete(&projectId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ProjectController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/project", uc.CreateProject)
	router.GET("/project/:id", uc.GetProject)
	router.GET("/project", uc.List)
	router.PATCH("/project", uc.UpdateProject)
	router.DELETE("/project/:id", uc.DeleteProject)
}

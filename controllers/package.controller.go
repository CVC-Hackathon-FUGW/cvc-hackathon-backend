package controllers

import (
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PackageController struct {
	PackageService services.PackageService
}

func NewPackageController(packageservice services.PackageService) PackageController {
	return PackageController{
		PackageService: packageservice,
	}
}

func (uc *PackageController) CreatePackage(ctx *gin.Context) {
	var pkg models.Package
	if err := ctx.ShouldBindJSON(&pkg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.PackageService.Create(&pkg)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *PackageController) GetPackage(ctx *gin.Context) {
	var pkgId string = ctx.Param("id")
	pkg, err := uc.PackageService.Show(&pkgId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pkg)
}

func (uc *PackageController) List(ctx *gin.Context) {
	pkgs, err := uc.PackageService.List()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pkgs)
}

func (uc *PackageController) UpdatePackage(ctx *gin.Context) {
	var pkg models.Package
	if err := ctx.ShouldBindJSON(&pkg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err := uc.PackageService.Update(&pkg)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pkg)
}

func (uc *PackageController) DeletePackage(ctx *gin.Context) {
	var pkgId string = ctx.Param("id")
	err := uc.PackageService.Delete(&pkgId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *PackageController) FindByAddress(ctx *gin.Context) {
	var address string = ctx.Param("address")
	pkgs, err := uc.PackageService.FindByAddress(&address)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pkgs)
}

func (uc *PackageController) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/package", uc.CreatePackage)
	r.GET("/package/:id", uc.GetPackage)
	r.GET("/package", uc.List)
	r.PATCH("/package", uc.UpdatePackage)
	r.DELETE("/package/:id", uc.DeletePackage)
	r.GET("/package/address/:address", uc.FindByAddress)
}

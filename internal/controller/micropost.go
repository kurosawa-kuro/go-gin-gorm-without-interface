package controller

import (
	"go-gin-gorm-without-interface/internal/models"
	services "go-gin-gorm-without-interface/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MicropostController struct {
	service *services.MicropostService
}

func NewMicropostController(service *services.MicropostService) *MicropostController {
	return &MicropostController{
		service: service,
	}
}

func (mc *MicropostController) Create(c *gin.Context) {
	var micropost models.Micropost
	if err := c.ShouldBindJSON(&micropost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mc.service.Create(&micropost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) GetAll(c *gin.Context) {
	microposts, err := mc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, microposts)
}

func (mc *MicropostController) GetByID(c *gin.Context) {
	micropost, err := mc.service.FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) Update(c *gin.Context) {
	micropost, err := mc.service.FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var newMicropost models.Micropost
	if err := c.ShouldBindJSON(&newMicropost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mc.service.Update(micropost, newMicropost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) Delete(c *gin.Context) {
	micropost, err := mc.service.FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := mc.service.Delete(micropost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

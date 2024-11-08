package controllers

import (
	"net/http"

	"go-gin-gorm-without-interface/internal/models"
	"go-gin-gorm-without-interface/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MicropostController struct {
	db *gorm.DB
}

func NewMicropostController(db *gorm.DB) *MicropostController {
	return &MicropostController{db: db}
}

func (mc *MicropostController) Create(c *gin.Context) {
	var micropost models.Micropost
	if err := c.ShouldBindJSON(&micropost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.NewMicropostService(mc.db).Create(&micropost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) GetAll(c *gin.Context) {
	microposts, err := services.NewMicropostService(mc.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, microposts)
}

func (mc *MicropostController) GetByID(c *gin.Context) {
	micropost, err := services.NewMicropostService(mc.db).FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) Update(c *gin.Context) {
	micropost, err := services.NewMicropostService(mc.db).FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var newMicropost models.Micropost
	if err := c.ShouldBindJSON(&newMicropost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.NewMicropostService(mc.db).Update(micropost, newMicropost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) Delete(c *gin.Context) {
	micropost, err := services.NewMicropostService(mc.db).FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := services.NewMicropostService(mc.db).Delete(micropost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

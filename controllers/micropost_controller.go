package controllers

import (
	"net/http"

	"go-gin-gorm-without-interface/models"

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

	if err := micropost.Create(mc.db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) GetAll(c *gin.Context) {
	microposts, err := models.GetAllMicroposts(mc.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, microposts)
}

func (mc *MicropostController) GetByID(c *gin.Context) {
	var micropost models.Micropost
	if err := micropost.FindByID(mc.db, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) Update(c *gin.Context) {
	var micropost models.Micropost
	if err := micropost.FindByID(mc.db, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var newMicropost models.Micropost
	if err := c.ShouldBindJSON(&newMicropost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := micropost.Update(mc.db, newMicropost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

func (mc *MicropostController) Delete(c *gin.Context) {
	var micropost models.Micropost
	if err := micropost.FindByID(mc.db, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := micropost.Delete(mc.db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

package controllers

import (
	"go-gin-gorm-without-interface/internal/models"
	"go-gin-gorm-without-interface/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title           Micropost API
// @version         1.0
// @description     This is a micropost server.
// @host           localhost:8080
// @BasePath       /

type MicropostController struct {
	service *services.MicropostService
}

func NewMicropostController(service *services.MicropostService) *MicropostController {
	return &MicropostController{
		service: service,
	}
}

// Create godoc
// @Summary      Create micropost
// @Description  Create a new micropost
// @Tags         microposts
// @Accept       json
// @Produce      json
// @Param        micropost  body      models.Micropost  true  "Micropost object"
// @Success      200       {object}  models.Micropost
// @Failure      400       {object}  object{error=string}
// @Failure      500       {object}  object{error=string}
// @Router       /microposts [post]
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

// GetAll godoc
// @Summary      List microposts
// @Description  Get all microposts
// @Tags         microposts
// @Produce      json
// @Success      200  {array}   models.Micropost
// @Failure      500  {object}  object{error=string}
// @Router       /microposts [get]
func (mc *MicropostController) GetAll(c *gin.Context) {
	microposts, err := mc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, microposts)
}

// GetByID godoc
// @Summary      Get micropost by ID
// @Description  Get a micropost by its ID
// @Tags         microposts
// @Produce      json
// @Param        id   path      string  true  "Micropost ID"
// @Success      200  {object}  models.Micropost
// @Failure      404  {object}  object{error=string}
// @Router       /microposts/{id} [get]
func (mc *MicropostController) GetByID(c *gin.Context) {
	micropost, err := mc.service.FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, micropost)
}

// Update godoc
// @Summary      Update micropost
// @Description  Update a micropost by its ID
// @Tags         microposts
// @Accept       json
// @Produce      json
// @Param        id        path      string           true  "Micropost ID"
// @Param        micropost body      models.Micropost true  "Micropost object"
// @Success      200       {object}  models.Micropost
// @Failure      400       {object}  object{error=string}
// @Failure      404       {object}  object{error=string}
// @Failure      500       {object}  object{error=string}
// @Router       /microposts/{id} [put]
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

// Delete godoc
// @Summary      Delete micropost
// @Description  Delete a micropost by its ID
// @Tags         microposts
// @Produce      json
// @Param        id   path      string  true  "Micropost ID"
// @Success      200  {object}  object{message=string}
// @Failure      404  {object}  object{error=string}
// @Failure      500  {object}  object{error=string}
// @Router       /microposts/{id} [delete]
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

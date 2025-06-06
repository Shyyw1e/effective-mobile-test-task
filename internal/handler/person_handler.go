package handler

import (
	"net/http"
	"strconv"

	"github.com/Shyyw1e/effective-mobile-test-task/internal/repository"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.EnrichService
	repo    *repository.PersonRepository
}

func NewHandler(repo *repository.PersonRepository, svc *service.EnrichService) *Handler {
	return &Handler{
		service: svc,
		repo:    repo,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/person", h.createPerson)
		api.GET("/person", h.getPeople)
		api.DELETE("/person/:id", h.deletePerson)
		api.PUT("/person/:id", h.updatePerson)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

type createPersonRequest struct {
	Name       string  `json:"name" binding:"required"`
	Surname    string  `json:"surname" binding:"required"`
	Patronymic *string `json:"patronymic"`
}

// @Summary Создать нового человека
// @Accept json
// @Produce json
// @Param input body createPersonRequest true "Информация о человеке"
// @Success 201 {object} model.Person
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /person [post]
func (h *Handler) createPerson(c *gin.Context) {
	var req createPersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	person, err := h.service.EnrichAndSave(req.Name, req.Surname, req.Patronymic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enrich and save"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (h *Handler) getPeople(c *gin.Context) {
	name := c.Query("name")
	gender := c.Query("gender")
	nationality := c.Query("nationality")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	people, err := h.repo.FindWithFilters(name, gender, nationality, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, people)
}

func (h *Handler) deletePerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.repo.DeleteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) updatePerson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Name       string  `json:"name" binding:"required"`
		Surname    string  `json:"surname" binding:"required"`
		Patronymic *string `json:"patronymic"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.repo.UpdateBasicInfo(uint(id), input.Name, input.Surname, input.Patronymic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

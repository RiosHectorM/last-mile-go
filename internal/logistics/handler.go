package logistics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreatePackage(c *gin.Context) {
	var p Package
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreatePackage(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (h *Handler) GetPackage(c *gin.Context) {
	id := c.Param("id") // Sacamos el ID de la URL
	p, err := h.service.GetPackage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paquete no encontrado"})
		return
	}

	c.JSON(http.StatusOK, p)
}

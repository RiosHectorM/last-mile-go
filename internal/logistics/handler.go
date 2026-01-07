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

func (h *Handler) GetAllPackages(c *gin.Context) {
	pkgs, err := h.service.GetAllPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista"})
		return
	}

	// Si la lista está vacía, devolvemos un array vacío [] en lugar de null
	if pkgs == nil {
		pkgs = []Package{}
	}

	c.JSON(http.StatusOK, pkgs)
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	if err := h.service.UpdatePackageStatus(id, body.Status); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Estado actualizado correctamente"})
}

func (h *Handler) DeletePackage(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeletePackage(id); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo eliminar el paquete"})
		return
	}
	c.JSON(200, gin.H{"message": "Paquete eliminado exitosamente"})
}
package logistics

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePackageHandler maneja la recepción de nuevos paquetes
func CreatePackageHandler(c *gin.Context) {
	var newPkg Package

	// Intentamos volcar el JSON que viene en la petición a nuestra estructura
	if err := c.ShouldBindJSON(&newPkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de envío inválidos"})
		return
	}

	// Simulamos que le asignamos un ID y fecha (esto luego lo hará la DB)
	newPkg.ID = "PKG-12345"
	newPkg.CreatedAt = time.Now()
	newPkg.Status = "pending"

	c.JSON(http.StatusCreated, gin.H{
		"message": "Paquete registrado en el sistema",
		"data":    newPkg,
	})
}

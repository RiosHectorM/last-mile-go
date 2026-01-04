package main

import (
	"github.com/RiosHectorM/last-mile-go/internal/logistics" // Ajustá a tu usuario de GitHub
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Ruta de salud
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Agrupamos las rutas de logística
	v1 := router.Group("/api/v1")
	{
		v1.POST("/packages", logistics.CreatePackageHandler)
	}

	router.Run(":8080")
}

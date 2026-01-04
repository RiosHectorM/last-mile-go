package main

import (
	"fmt"

	"github.com/RiosHectorM/last-mile-go/internal/database"
	"github.com/RiosHectorM/last-mile-go/internal/logistics" // Ajustá a tu usuario de GitHub
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.GetConnection()
	if err != nil {
		panic(fmt.Sprintf("No se pudo conectar a la DB: %v", err))
	}
	database.Migraciones(db)
	defer db.Close()
	router := gin.Default()

	// Ruta de salud
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Agrupamos las rutas de logística
	v1 := router.Group("/api/v1")
	{
		v1.POST("/packages", logistics.CreatePackageHandler(db))
	}

	router.Run(":8080")
}

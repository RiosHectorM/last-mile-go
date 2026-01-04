package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inicializamos el motor de Gin
	router := gin.Default()

	// 2. Definimos una ruta básica de prueba
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "¡Sistema de logística activo!",
			"status":  "ok",
		})
	})

	// 3. Arrancamos el servidor en el puerto 8080
	router.Run(":8080")
}

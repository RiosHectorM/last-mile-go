package main

import (
	"fmt"
	"log"

	"github.com/RiosHectorM/last-mile-go/internal/database"
	"github.com/RiosHectorM/last-mile-go/internal/logistics" // Ajustá a tu usuario de GitHub
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.GetConnection()
	if err != nil {
		panic(fmt.Sprintf("No se pudo conectar a la DB: %v", err))
	}

	defer db.Close()

	if err := database.Migraciones(db); err != nil {
		log.Fatalf("Error en migraciones: %v", err)
	}

	repo := logistics.NewRepository(db)
	service := logistics.NewService(repo)
	handler := logistics.NewHandler(service)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/packages", handler.CreatePackage)
		api.GET("/packages/:id", handler.GetPackage)
		api.GET("/packages", handler.GetAllPackages)
	}

	fmt.Println("🚀 Servidor corriendo en el puerto 8080")
	router.Run(":8080")
}

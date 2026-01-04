package logistics

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// CreatePackageHandler maneja la recepción de nuevos paquetes
func CreatePackageHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p Package
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		query := `INSERT INTO packages (tracking_code, receiver_name, destination, weight, status) 
				  VALUES ($1, $2, $3, $4, $5) RETURNING id`

		err := db.QueryRow(query, p.TrackingCode, p.ReceiverName, p.Destination, p.Weight, "pending").Scan(&p.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al guardar en DB: " + err.Error()})
			return
		}

		c.JSON(201, p)
	}
}

package logistics

import "time"

// Package representa un envío
type Package struct {
	ID           string    `json:"id"`
	TrackingCode string    `json:"tracking_code"`
	ReceiverName string    `json:"receiver_name" binding:"required"`
	Destination  string    `json:"destination" binding:"required"`
	Weight       float64   `json:"weight" binding:"required"`
	Status       string    `json:"status"` // "pending", "in_transit", "delivered"
	CreatedAt    time.Time `json:"created_at"`
}

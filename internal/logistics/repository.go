package logistics

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(p *Package) error {
	query := `INSERT INTO packages (tracking_code, receiver_name, destination, weight, status) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	return r.db.QueryRow(query, p.TrackingCode, p.ReceiverName, p.Destination, p.Weight, p.Status).
		Scan(&p.ID, &p.CreatedAt)
}

func (r *Repository) GetByID(id string) (*Package, error) {
	query := `SELECT id, tracking_code, receiver_name, destination, weight, status, created_at 
			  FROM packages WHERE id = $1`

	p := &Package{}
	err := r.db.QueryRow(query, id).Scan(
		&p.ID, &p.TrackingCode, &p.ReceiverName, &p.Destination, &p.Weight, &p.Status, &p.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

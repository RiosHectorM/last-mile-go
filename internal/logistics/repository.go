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

func (r *Repository) GetAll() ([]Package, error) {
	query := `SELECT id, tracking_code, receiver_name, destination, weight, status, created_at FROM packages`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []Package
	for rows.Next() {
		var p Package
		if err := rows.Scan(&p.ID, &p.TrackingCode, &p.ReceiverName, &p.Destination, &p.Weight, &p.Status, &p.CreatedAt); err != nil {
			return nil, err
		}
		packages = append(packages, p)
	}
	return packages, nil
}

func (r *Repository) UpdateStatus(id string, status string) error {
	query := `UPDATE packages SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}
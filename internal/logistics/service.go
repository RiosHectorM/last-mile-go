package logistics

import (
	"fmt"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePackage(p *Package) error {
	// REGLA DE NEGOCIO: Generar un Tracking Code automático si no viene uno
	if p.TrackingCode == "" {
		p.TrackingCode = "TRK-" + fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// REGLA DE NEGOCIO: Validar peso mínimo
	if p.Weight <= 0 {
		return fmt.Errorf("el peso del paquete debe ser mayor a 0")
	}

	p.Status = "pending"
	return s.repo.Save(p)
}

func (s *Service) GetPackage(id string) (*Package, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllPackages() ([]Package, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdatePackageStatus(id string, status string) error {
	// Validación básica de estados de logística
	valid := false
	for _, s := range []string{"pending", "in_transit", "delivered", "canceled"} {
		if s == status {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("estado no válido: %s", status)
	}

	return s.repo.UpdateStatus(id, status)
}

func (s *Service) DeletePackage(id string) error {
	return s.repo.Delete(id)
}
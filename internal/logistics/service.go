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

func (s *Service) DeletePackage(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) UpdatePackageStatus(id string, newStatus string) error {
	// 1. Buscamos el paquete actual para saber en qué estado está
	currentPkg, err := s.repo.GetByID(id)
	if err != nil {
		return err // El repo ya devuelve "not found" si no existe
	}

	// 2. Validamos la lógica de negocio
	switch currentPkg.Status {
	case StatusPending:
		if newStatus != StatusInTransit {
			return fmt.Errorf("un paquete pendiente solo puede pasar a 'in_transit'")
		}
	case StatusInTransit:
		if newStatus != StatusDelivered {
			return fmt.Errorf("un paquete en tránsito solo puede pasar a 'delivered'")
		}
	case StatusDelivered:
		return fmt.Errorf("el paquete ya fue entregado y no puede cambiar de estado")
	default:
		// Por si viene un estado que no conocemos
		if newStatus != StatusPending && newStatus != StatusInTransit && newStatus != StatusDelivered {
			return fmt.Errorf("estado no válido: %s", newStatus)
		}
	}

	// 3. Si pasó las reglas, recién ahí llamamos al repo
	return s.repo.UpdateStatus(id, newStatus)
}

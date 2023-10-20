package service

import (
	"errors"

	"github.com/gustavohmsilva/ports/internal/core/domain"
	"github.com/gustavohmsilva/ports/internal/core/port"
)

type HarborServiceIface interface {
	CreateHarbor(harbor domain.Harbor) (string, error)
	UpdateHarbor(harbor domain.Harbor) error
}

// Harbor is the service structure, contains all business logic operations in
// its methods
type Harbor struct {
	repository port.HarborRepositoryIface
}

// NewHarbor instantiate a new harbor service with the repository that will
// be used for persistence
func NewHarbor(repository port.HarborRepositoryIface) port.HarborServiceIface {
	return &Harbor{
		repository: repository,
	}
}

// CreateHarbor will do all the necessary business logic validations prior
// to forwarding the object for the corresponding database operation (create)
func (h *Harbor) CreateHarbor(harbor domain.Harbor) (string, error) {
	if harbor.InternalID != "" {
		harbor.InternalID = ""
	}
	return h.repository.CreateHarbor(harbor)
}

// UpdateHarbor will do all the necessary business logic validations prior
// to forwarding the object for the corresponding database operation (update)
func (h *Harbor) UpdateHarbor(harbor domain.Harbor) error {
	if harbor.InternalID == "" {
		return errors.New("harbor don't contain uuid")
	}
	return h.repository.UpdateHarbor(harbor)
}

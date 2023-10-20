package port

import (
	"net/http"

	"github.com/gustavohmsilva/ports/internal/core/domain"
)

// mockgen --source=internal/core/port/harbor.go --destination=test/mocks/harbor_port.go

// HarborRepositoryIface is the connection protocol from domain to repository
type HarborRepositoryIface interface {
	CreateHarbor(harbor domain.Harbor) (string, error)
	UpdateHarbor(harbor domain.Harbor) error
}

// HarborServiceIface is the connection protocol from service to domain
type HarborServiceIface interface {
	CreateHarbor(harbor domain.Harbor) (string, error)
	UpdateHarbor(harbor domain.Harbor) error
}

// HarborHandlerIface is the connection protocol from handler to domain
type HarborHandlerIface interface {
	CreateHarbor(w http.ResponseWriter, r *http.Request)
	UpdateHarbor(w http.ResponseWriter, r *http.Request)
}

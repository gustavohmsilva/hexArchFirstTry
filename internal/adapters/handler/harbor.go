package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gustavohmsilva/ports/internal/core/domain"
	"github.com/gustavohmsilva/ports/internal/core/port"
)

type HarborHandlerIface interface {
	CreateHarbor(w http.ResponseWriter, r *http.Request)
	UpdateHarbor(w http.ResponseWriter, r *http.Request)
}

// Harbor is the handler relative to dealing with all harbor CRUD operations
type Harbor struct {
	service port.HarborServiceIface
}

// NewHarbor instantiate a new Harbor handler
func NewHarbor(portSvc port.HarborServiceIface) port.HarborHandlerIface {
	return &Harbor{
		service: portSvc,
	}
}

// CreateHarbor is the handler responsible for dealing with the creation of new
// Harbors in the microservice database
func (h *Harbor) CreateHarbor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var harbor domain.Harbor
	if err := json.NewDecoder(r.Body).Decode(&harbor); err != nil {
		log.Printf("failed to unmarshal request body: %v", err)
		w = writeErrorResponse(w, http.StatusBadRequest, "failed to read harbor")
		return
	}

	assignedID, err := h.service.CreateHarbor(harbor)
	if err != nil {
		log.Printf("failed to create harbor in database: %v", err)
		w = writeErrorResponse(w, http.StatusInternalServerError, "failed to save harbor")
		return
	}
	w = writeSuccessResponse(w, http.StatusCreated, assignedID)

}

// UpdateHarbor is the handler responsible for dealing with the update of
// existing harbors in the microservice database.
func (h *Harbor) UpdateHarbor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var harbor domain.Harbor
	if err := json.NewDecoder(r.Body).Decode(&harbor); err != nil {
		log.Printf("failed to unmarshal request body: %v", err)
		w = writeErrorResponse(w, http.StatusBadRequest, "failed to read harbor")
		return
	}

	if err := h.service.UpdateHarbor(harbor); err != nil {
		log.Printf("failed to update harbor in database: %v", err)
		w = writeErrorResponse(w, http.StatusInternalServerError, "failed to save harbor")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

package repository

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/gustavohmsilva/ports/internal/core/domain"
)

// HarborRepositoryIface is the connection protocol from domain to repository
type HarborRepositoryIface interface {
	CreateHarbor(harbor domain.Harbor) (string, error)
	UpdateHarbor(harbor domain.Harbor) error
}

// Harbor is the repository relative to all harbor CRUD operations at the
// in memory database
type Harbor struct {
	DB map[string]domain.Harbor
	sync.Mutex
}

// NewHarbor instantiate a new Harbor repository
func NewHarbor() HarborRepositoryIface {
	return &Harbor{
		DB: make(map[string]domain.Harbor),
	}
}

// CreateHarbor stores in the in-memory database a new harbor, it also gives
// it an internal UUID
func (h *Harbor) CreateHarbor(harbor domain.Harbor) (string, error) {
	h.Lock()
	defer h.Unlock()
	internalID := uuid.New().String()
	h.DB[internalID] = harbor
	h.echoStoredHarbours()
	return internalID, nil
}

// UpdateHarbor updates an existing harbor in the in-memory database. It will
// fail if it can't locate a previously existing one. No validation of field
// change is made as this is a simplified project for recruitment purposes
func (h *Harbor) UpdateHarbor(harbor domain.Harbor) error {
	h.Lock()
	defer h.Unlock()

	_, ok := h.DB[harbor.InternalID]
	if !ok {
		log.Print("received harbor isn't present in the database")
		return errors.New("can't save harbor in the database")
	}
	h.DB[harbor.InternalID] = harbor
	h.echoStoredHarbours()
	return nil
}

func (h *Harbor) echoStoredHarbours() {
	t := table.NewWriter()
	t.SetCaption("HARBORS")
	t.AppendHeader(table.Row{"Internal ID", "External ID", "Name", "City", "Country", "Alias", "Region", "Latitude", "Longitude", "Province", "Timezone", "Unlocs", "Code"})
	for k, v := range h.DB {
		t.AppendRow(table.Row{k, v.ExternalID, v.Name, v.City, v.Country, v.Alias, v.Regions, v.Coordinates[0], v.Coordinates[1], v.Province, v.Timezone, v.Unlocs, v.Code})
		fmt.Println("test")
	}
	fmt.Println(t.Render())
}

package service

import (
	"github.com/gustavohmsilva/ports/internal/core/domain"
)

func testData_AEAJM() domain.Harbor {
	return domain.Harbor{
		ExternalID:  "AEAJM",
		Name:        "Ajman",
		City:        "United Arab Emirates",
		Country:     "United Arab Emirates",
		Coordinates: []float32{55.513645, 25.405216},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}
}

func testData_AEAJM_withUUID() domain.Harbor {
	return domain.Harbor{
		InternalID:  "c7f8b388-6e9a-11ee-b962-0242ac120002",
		ExternalID:  "AEAJM",
		Name:        "Ajman",
		City:        "United Arab Emirates",
		Country:     "United Arab Emirates",
		Coordinates: []float32{55.513645, 25.405216},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}
}
func testData_UUID() string {
	return "c7f8b388-6e9a-11ee-b962-0242ac120002"
}

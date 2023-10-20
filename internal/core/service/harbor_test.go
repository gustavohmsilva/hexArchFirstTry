package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_port "github.com/gustavohmsilva/ports/test/mocks"
)

func TestHarbor_CreateHarbor(t *testing.T) {
	t.Run("Create simple Harbor", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		harborRepositoryMock := mock_port.NewMockHarborRepositoryIface(ctrl)
		harborRepositoryMock.EXPECT().CreateHarbor(testData_AEAJM()).Return(testData_UUID(), nil).Times(1)
		nh := NewHarbor(harborRepositoryMock)
		id, err := nh.CreateHarbor(testData_AEAJM())
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})
	t.Run("Database fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		harborRepositoryMock := mock_port.NewMockHarborRepositoryIface(ctrl)
		harborRepositoryMock.EXPECT().CreateHarbor(testData_AEAJM()).Return("", errors.New("sample error")).Times(1)
		nh := NewHarbor(harborRepositoryMock)
		id, err := nh.CreateHarbor(testData_AEAJM())
		assert.Error(t, err)
		assert.Empty(t, id)
	})
}

func TestHarbor_UpdateHarbor(t *testing.T) {
	t.Run("Create update Harbor", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		harborRepositoryMock := mock_port.NewMockHarborRepositoryIface(ctrl)
		harborRepositoryMock.EXPECT().UpdateHarbor(testData_AEAJM_withUUID()).Return(nil).Times(1)
		nh := NewHarbor(harborRepositoryMock)
		err := nh.UpdateHarbor(testData_AEAJM_withUUID())
		assert.NoError(t, err)
	})
	t.Run("Database fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		harborRepositoryMock := mock_port.NewMockHarborRepositoryIface(ctrl)
		nh := NewHarbor(harborRepositoryMock)
		err := nh.UpdateHarbor(testData_AEAJM())
		assert.Error(t, err)
	})
}

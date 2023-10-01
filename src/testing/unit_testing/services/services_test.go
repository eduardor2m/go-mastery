package services

import (
	"testing"

	"github.com/eduardor2m/go-mastery/src/testing/unit_testing/mocks"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockRepository(ctrl)

	mockRepository.EXPECT().Create("Dudu").Return(nil)

	service := NewService(mockRepository)

	err := service.Create("Dudu")

	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockRepository(ctrl)

	mockRepository.EXPECT().GetAll().Return([]string{"Dudu"})

	service := NewService(mockRepository)

	result := service.GetAll()

	if len(result) != 1 {
		t.Errorf("Expected 1, got %d", len(result))
	}
}

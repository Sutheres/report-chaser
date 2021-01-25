package service

import (
	"github.com/Sutheres/report-chaser/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewServiceMock(t *testing.T)  {
	mockCtl := gomock.NewController(t)
	secMock := mocks.NewMockSEC(mockCtl)

	svc := NewService(
		"Test", "Test",
		WithSEC(secMock),
	)

	if svc == nil {
		t.Error("Error: uninitialized")
	}
}
package service

import (
	"productive-go-client/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Setup mock
type MockProductive struct {
	mock.Mock
}

func (m *MockProductive) PostTimeEntry(config models.Config, timeEntry models.TimeEntry) error {
	args := m.Called(config, timeEntry)
	return args.Error(0)
}

func (m *MockProductive) GetServiceAssignments(config models.Config, date string) ([]models.ServiceAssignment, error) {
	args := m.Called(config)
	return args.Get(0).([]models.ServiceAssignment), args.Error(1)
}

func (m *MockProductive) GetUser(config models.Config) (models.User, error) {
	args := m.Called(config)
	return args.Get(0).(models.User), args.Error(1)
}

// Test
func TestEnterTime(t *testing.T) {
	mockRepo := new(MockProductive)
	timeService := NewTimeService(mockRepo)

	user := &models.User{ID: "1"}
	config := &models.Config{AccessToken: "token", CompanyId: "company", UserEmail: "user@example.com"}
	date := "2021-01-01"
	serviceAssignmentID := "serviceAssignmentID"
	notes := "Test note"
	timeHours := 2
	timeMinutes := 30

	mockRepo.On("PostTimeEntry", *config, mock.AnythingOfType("models.TimeEntry")).Return(nil)

	err := timeService.EnterTime(user, config, date, serviceAssignmentID, notes, timeHours, timeMinutes)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

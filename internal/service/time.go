package service

import (
	"productive-go-client/internal/data"
	"productive-go-client/internal/models"
)

// TimeService is defined within the service package and wraps the ProductiveService interface.
type TimeService struct {
	Repo data.ProductiveService // Use the interface type directly, without a pointer
}

// NewTimeService constructs a new TimeService with the given ProductiveService implementation.
func NewTimeService(repo data.ProductiveService) *TimeService { // Accept the interface type
	return &TimeService{Repo: repo}
}

// EnterTime encapsulates the business logic for creating and posting a time entry.
func (s *TimeService) EnterTime(user *models.User, config *models.Config, date string, serviceAssignmentID string, notes string, timeHours, timeMinutes int) error {
	// Construct the time entry
	timeEntry := models.NewTimeEntry()
	timeEntry.Attributes.Date = date
	timeEntry.Relationships.Service.Data.ID = serviceAssignmentID
	timeEntry.Attributes.Note = notes
	timeEntry.Attributes.Time = timeMinutes + timeHours*60
	timeEntry.Relationships.Person.Data.ID = user.ID

	// Post the time entry using the ProductiveService interface
	return s.Repo.PostTimeEntry(*config, timeEntry) // Ensure timeEntry is correctly referenced or passed as per your actual method signature
}

func (s *TimeService) GetServiceAssignments(config models.Config) ([]models.ServiceAssignment, error) {
	return s.Repo.GetServiceAssignments(config)
}

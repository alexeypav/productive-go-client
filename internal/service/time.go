package service

import (
	"productive-go-client/internal/data"
	"productive-go-client/internal/models"
)

// Assuming TimeService is defined within the service package and wraps the repository.
type TimeService struct {
	Repo *data.Productive
}

func NewTimeService(repo *data.Productive) *TimeService {
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

	// Post the time entry
	return s.Repo.PostTimeEntry(*config, timeEntry)
}

func (s *TimeService) GetServiceAssignments(config models.Config) ([]models.ServiceAssignment, error) {
	return s.Repo.GetServiceAssignments(config)
}

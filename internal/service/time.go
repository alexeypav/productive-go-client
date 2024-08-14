package service

import (
	"productive-go-client/internal/data"
	"productive-go-client/internal/models"
)

type TimeService struct {
	Repo data.ProductiveService
}

func NewTimeService(repo data.ProductiveService) *TimeService {
	return &TimeService{Repo: repo}
}

func (s *TimeService) EnterTime(user *models.User, config *models.Config, date string, serviceAssignmentID string, notes string, timeHours, timeMinutes int) error {
	// Construct the time entry
	timeEntry := models.NewTimeEntry()
	timeEntry.Attributes.Date = date
	timeEntry.Relationships.Service.Data.ID = serviceAssignmentID
	timeEntry.Attributes.Note = notes
	timeEntry.Attributes.Time = timeMinutes + timeHours*60
	timeEntry.Relationships.Person.Data.ID = user.ID

	// Post the time entry using the ProductiveService interface
	return s.Repo.PostTimeEntry(*config, timeEntry)
}

func (s *TimeService) GetServiceAssignments(config models.Config, date string) ([]models.ServiceAssignment, error) {
	return s.Repo.GetServiceAssignments(config, date)
}

package models

// Broken up for easier comprehension
type TimeEntry struct {
	Type          string         `json:"type"`
	Attributes    TimeAttributes `json:"attributes"`
	Relationships Relationships  `json:"relationships"`
}

type TimeAttributes struct {
	Note string `json:"note"`
	Date string `json:"date"`
	Time int    `json:"time"`
}

type Relationships struct {
	Person  Relationship `json:"person"`
	Service Relationship `json:"service"`
}

type Relationship struct {
	Data RelationshipData `json:"data"`
}

type RelationshipData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// NewTimeEntry creates a new TimeEntry with default values, only this should be used as the time entry body
func NewTimeEntry() TimeEntry {
	return TimeEntry{
		Type: "time_entries",
		Attributes: TimeAttributes{
			Note: "",
			Date: "2022-01-01",
			Time: 60,
		},
		Relationships: Relationships{
			Person: Relationship{
				Data: RelationshipData{
					Type: "people",
					ID:   "defaultPersonId",
				},
			},
			Service: Relationship{
				Data: RelationshipData{
					Type: "services",
					ID:   "defaultServiceId",
				},
			},
		},
	}
}

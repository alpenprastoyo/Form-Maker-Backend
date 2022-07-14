package survey

import "time"

type Survey struct {
	ID          int
	NameSurvey  string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

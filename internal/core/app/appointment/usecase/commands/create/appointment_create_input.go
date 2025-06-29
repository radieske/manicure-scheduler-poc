package appointments

import "time"

type CreateAppointmentInput struct {
	Name        string
	Phone       string
	Service     string
	ScheduledAt time.Time
}

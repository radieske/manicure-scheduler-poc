package appointment

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID          string
	Name        string
	Phone       string
	Service     string
	ScheduledAt time.Time
	CreatedAt   time.Time
}

func Create(name, phone, service string, scheduledAt time.Time) Appointment {
	return Appointment{
		ID:          uuid.New().String(),
		Name:        name,
		Phone:       phone,
		Service:     service,
		ScheduledAt: scheduledAt,
		CreatedAt:   time.Now(),
	}
}

func With(id, name, phone, service string, scheduledAt, createdAt time.Time) Appointment {
	return Appointment{
		ID:          id,
		Name:        name,
		Phone:       phone,
		Service:     service,
		ScheduledAt: scheduledAt,
		CreatedAt:   createdAt,
	}
}

package request

import (
	"time"

	create_appointments "github.com/radieske/manicure-scheduler-poc/internal/core/app/appointment/usecase/commands/create"
)

type CreateAppointmentDTO struct {
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Service     string    `json:"service"`
	ScheduledAt time.Time `json:"scheduled_at"`
}

func (dto CreateAppointmentDTO) ToInput() create_appointments.CreateAppointmentInput {
	return create_appointments.CreateAppointmentInput{
		Name:        dto.Name,
		Phone:       dto.Phone,
		Service:     dto.Service,
		ScheduledAt: dto.ScheduledAt,
	}
}

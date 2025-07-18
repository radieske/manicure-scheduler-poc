package response

import (
	"time"

	"github.com/radieske/manicure-scheduler-poc/internal/core/domain/appointment"
)

type AppointmentResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Service     string    `json:"service"`
	ScheduledAt time.Time `json:"scheduledAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

func FromEntity(app appointment.Appointment) AppointmentResponse {
	return AppointmentResponse{
		ID:          app.ID,
		Name:        app.Name,
		Phone:       app.Phone,
		Service:     app.Service,
		ScheduledAt: app.ScheduledAt,
		CreatedAt:   app.CreatedAt,
	}
}

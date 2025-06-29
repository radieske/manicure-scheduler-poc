package appointments

import (
	"context"

	"github.com/radieske/manicure-scheduler-poc/internal/core/domain/appointment"
)

type CreateAppointmentUseCase interface {
	Execute(ctx context.Context, input CreateAppointmentInput) (appointment.Appointment, error)
}

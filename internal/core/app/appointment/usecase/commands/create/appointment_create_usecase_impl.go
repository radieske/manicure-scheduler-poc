package appointments

import (
	"context"
	"errors"

	"github.com/radieske/manicure-scheduler-poc/internal/core/domain/appointment"
)

type createAppointmentUseCaseImpl struct {
	ctx  context.Context
	repo appointment.Repository
}

func NewCreateAppointmentUseCase(
	ctx context.Context,
	repo appointment.Repository,
) CreateAppointmentUseCase {
	return &createAppointmentUseCaseImpl{
		ctx:  ctx,
		repo: repo,
	}
}

func (uc *createAppointmentUseCaseImpl) Execute(ctx context.Context, input CreateAppointmentInput) (appointment.Appointment, error) {
	if input.Name == "" || input.Phone == "" || input.Service == "" {
		return appointment.Appointment{}, errors.New("É necessário informar nome, telefone e serviço")
	}

	appt := appointment.Create(
		input.Name,
		input.Phone,
		input.Service,
		input.ScheduledAt,
	)

	if err := uc.repo.Create(ctx, appt); err != nil {
		return appointment.Appointment{}, err
	}

	return appt, nil
}

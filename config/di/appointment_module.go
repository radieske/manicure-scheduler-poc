package di

import (
	"context"

	"gorm.io/gorm"

	create_appointments "github.com/radieske/manicure-scheduler-poc/internal/core/app/appointment/usecase/commands/create"
	appointmentdb "github.com/radieske/manicure-scheduler-poc/internal/infra/db/sql/pg/gorm/appointments"
)

// AppointmentModule agrupa todas as dependências de Appointment
type AppointmentModule struct {
	CreateAppointmentUseCase create_appointments.CreateAppointmentUseCase
}

// NewAppointmentModule instancia e injeta todas as dependências do módulo de agendamento
func NewAppointmentModule(ctx context.Context, db *gorm.DB) *AppointmentModule {
	repo := appointmentdb.NewRepository(ctx, db)
	usecase := create_appointments.NewCreateAppointmentUseCase(ctx, repo)

	return &AppointmentModule{
		CreateAppointmentUseCase: usecase,
	}
}

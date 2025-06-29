package appointmentdb

import (
	"context"

	"github.com/radieske/manicure-scheduler-poc/internal/core/domain/appointment"
	"gorm.io/gorm"
)

type Repository struct {
	ctx context.Context
	db  *gorm.DB
}

func NewRepository(
	ctx context.Context,
	db *gorm.DB,
) *Repository {
	return &Repository{
		ctx: ctx,
		db:  db,
	}
}

func (r *Repository) Create(ctx context.Context, a appointment.Appointment) error {
	entity := AppointmentEntity{
		ID:          a.ID,
		Name:        a.Name,
		Phone:       a.Phone,
		Service:     a.Service,
		ScheduledAt: a.ScheduledAt,
		CreatedAt:   a.CreatedAt,
	}
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *Repository) FindByPhone(ctx context.Context, phone string) ([]appointment.Appointment, error) {
	var entities []AppointmentEntity
	if err := r.db.WithContext(ctx).Where("phone = ?", phone).Find(&entities).Error; err != nil {
		return nil, err
	}

	var result []appointment.Appointment
	for _, e := range entities {
		result = append(result, appointment.With(
			e.ID,
			e.Name,
			e.Phone,
			e.Service,
			e.ScheduledAt,
			e.CreatedAt,
		))
	}
	return result, nil
}

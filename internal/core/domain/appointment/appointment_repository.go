package appointment

import "context"

type Repository interface {
	Create(ctx context.Context, a Appointment) error
	FindByPhone(ctx context.Context, phone string) ([]Appointment, error)
}

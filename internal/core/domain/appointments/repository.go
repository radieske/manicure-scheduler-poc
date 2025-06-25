package appointment

type Repository interface {
	Create(appointment Appointment) error
	FindByPhone(phone string) ([]Appointment, error)
}

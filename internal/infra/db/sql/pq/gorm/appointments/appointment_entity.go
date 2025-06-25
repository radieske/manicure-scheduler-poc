package appointmentdb

import (
	"time"
)

type AppointmentEntity struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"type:text;not null"`
	Phone       string    `gorm:"type:text;not null"`
	Service     string    `gorm:"type:text;not null"`
	ScheduledAt time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (AppointmentEntity) TableName() string {
	return "appointments"
}

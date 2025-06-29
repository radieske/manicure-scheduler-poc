package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/radieske/manicure-scheduler-poc/config/di"
	"github.com/radieske/manicure-scheduler-poc/internal/infra/web/http/handler/appointment"
	"gorm.io/gorm"
)

func New(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	ctx := context.Background()
	appModule := di.NewAppointmentModule(ctx, db)
	handler := appointment.NewHandler(appModule.CreateAppointmentUseCase)

	r.Route("/appointments", func(r chi.Router) {
		r.Post("/", handler.CreateAppointment)
	})

	return r
}

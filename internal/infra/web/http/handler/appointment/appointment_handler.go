package appointment

import (
	"encoding/json"
	"net/http"

	create_appointments "github.com/radieske/manicure-scheduler-poc/internal/core/app/appointment/usecase/commands/create"
	"github.com/radieske/manicure-scheduler-poc/internal/infra/web/http/handler/appointment/request"
	"github.com/radieske/manicure-scheduler-poc/internal/infra/web/http/handler/appointment/response"
)

type Handler struct {
	CreateAppointmentUseCase create_appointments.CreateAppointmentUseCase
}

func NewHandler(createAppointmentUseCase create_appointments.CreateAppointmentUseCase) *Handler {
	return &Handler{
		CreateAppointmentUseCase: createAppointmentUseCase,
	}
}

func (h *Handler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var inputDTO request.CreateAppointmentDTO

	if err := json.NewDecoder(r.Body).Decode(&inputDTO); err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	app, err := h.CreateAppointmentUseCase.Execute(r.Context(), inputDTO.ToInput())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := response.FromEntity(app)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

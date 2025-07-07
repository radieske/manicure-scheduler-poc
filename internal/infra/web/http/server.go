package server

import (
	"log"
	"net/http"

	gorm_config "github.com/radieske/manicure-scheduler-poc/internal/infra/db/sql/pg/gorm"
	"github.com/radieske/manicure-scheduler-poc/internal/infra/web/http/middleware"
	"github.com/radieske/manicure-scheduler-poc/internal/infra/web/http/router"
)

func Start() error {
	db := gorm_config.ConnectDB()
	r := router.New(db)

	// Aplica o middleware CORS
	handlerWithCORS := middleware.CORSMiddleware(r)

	log.Println("Server running at http://localhost:3333")
	return http.ListenAndServe(":3333", handlerWithCORS)
}

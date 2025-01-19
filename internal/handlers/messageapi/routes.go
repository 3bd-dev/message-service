package messageapi

import (
	"net/http"

	"github.com/3bd-dev/wallet-service/internal/services/message"
	"github.com/gorilla/mux"
)

type Config struct {
	Service *message.Service
}

// Routes adds specific routes for this group.
func Routes(router *mux.Router, cfg Config) {
	api := newapi(cfg.Service)
	messages := router.PathPrefix("/api/v1/messages").Subrouter()
	messages.HandleFunc("", api.query).Methods(http.MethodGet)
	messages.HandleFunc("", api.create).Methods(http.MethodPost)
	messages.HandleFunc("/{id}/reviews", api.review).Methods(http.MethodPost)
}

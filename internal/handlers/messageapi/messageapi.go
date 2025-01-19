package messageapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/3bd-dev/wallet-service/internal/services/message"
	"github.com/3bd-dev/wallet-service/pkg/errs"
	"github.com/3bd-dev/wallet-service/pkg/page"
	"github.com/3bd-dev/wallet-service/pkg/query"
	"github.com/3bd-dev/wallet-service/pkg/web"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type api struct {
	service *message.Service
}

func newapi(svc *message.Service) *api {
	return &api{service: svc}
}

// create creates a new wallet.
func (a *api) create(w http.ResponseWriter, r *http.Request) {
	var req message.NewMessage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		web.RenderErr(w, errs.New(errs.InvalidArgument, fmt.Errorf("failed to decode request body: %w", err)))
		return
	}

	wallet, err := a.service.Create(r.Context(), &req)
	if err != nil {
		web.RenderErr(w, err)
		return
	}

	web.RenderOk(w, wallet)
}

func (a *api) review(w http.ResponseWriter, r *http.Request) {
	var req message.NewMessageReview

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	messageID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		web.RenderErr(w, errs.New(errs.InvalidArgument, err))
		return
	}

	message, err := a.service.Review(r.Context(), messageID, &req)
	if err != nil {
		web.RenderErr(w, err)
		return
	}

	web.RenderOk(w, message)
}

func (a *api) query(w http.ResponseWriter, r *http.Request) {
	qp := parseQueryParams(r)

	page, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		web.RenderErr(w, errs.NewFieldErrors("page", err))
		return
	}

	filter, err := parseFilter(qp)
	if err != nil {
		web.RenderErr(w, err.(*errs.Error))
		return
	}

	msgs, total, err := a.service.Query(r.Context(), filter, page)
	if err != nil {
		web.RenderErr(w, err)
		return
	}

	web.RenderOk(w, query.NewResult(msgs, total, page))
}

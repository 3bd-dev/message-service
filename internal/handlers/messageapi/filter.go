package messageapi

import (
	"net/http"

	"github.com/3bd-dev/wallet-service/internal/services/message"
	"github.com/3bd-dev/wallet-service/pkg/errs"
	"github.com/google/uuid"
)

type queryParams struct {
	SenderID string `json:"sender_id"`
	Page     string `json:"page"`
	Rows     string `json:"rows"`
}

func parseQueryParams(r *http.Request) queryParams {
	values := r.URL.Query()

	filter := queryParams{
		Page:     values.Get("page"),
		Rows:     values.Get("row"),
		SenderID: values.Get("sender_id"),
	}

	return filter
}

func parseFilter(qp queryParams) (message.QueryFilter, error) {
	var fieldErrors errs.FieldErrors
	var filter message.QueryFilter

	if qp.SenderID != "" {
		id, err := uuid.Parse(qp.SenderID)
		switch err {
		case nil:
			filter.SenderID = &id
		default:
			fieldErrors.Add("sender_id", err)
		}
	}

	if fieldErrors != nil {
		return message.QueryFilter{}, fieldErrors.ToError()
	}

	return filter, nil
}

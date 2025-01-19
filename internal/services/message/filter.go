package message

import (
	"github.com/google/uuid"
)

type QueryFilter struct {
	SenderID *uuid.UUID
}

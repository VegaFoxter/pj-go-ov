package domain

import (
	"github.com/google/uuid"
)

type Session struct {
	UserId uint64
	UUID   uuid.UUID
}

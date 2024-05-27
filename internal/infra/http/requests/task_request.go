package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline" validate:"required"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    time.Unix(r.Deadline, 0),
	}, nil
}

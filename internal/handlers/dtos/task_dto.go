package dtos

import (
	"github.com/google/uuid"
)

type CreateTaskRequestDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequestDto struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
}

func (utr *UpdateTaskRequestDto) ValidTitleField() bool {
	return utr.Title != ""
}

func (ctr *CreateTaskRequestDto) ValidTitleField() bool {
	return ctr.Title != ""
}

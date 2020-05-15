package domain

import "github.com/google/uuid"

type Repository interface {
	Function(id uuid.UUID) error
}

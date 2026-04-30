package postgres

import "github.com/google/uuid"

type actionRepositoryEntity struct {
	ID          uuid.UUID
	Name        string
	Description string
	Config      map[string]string
	Type        string
}

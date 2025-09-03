package ltrc

import "github.com/google/uuid"

type Event struct {
	ID     uuid.UUID
	Format Format
}

type Racer struct {
	Event uuid.UUID
	Name  string
	Mmr   float64
	Score float64
}

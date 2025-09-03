package ltrc

import "github.com/google/uuid"

type Event struct {
	Event  uuid.UUID
	Format Format
	Racers Racer
}

type Racer struct {
	Name  string
	Mmr   float64
	Score float64
}

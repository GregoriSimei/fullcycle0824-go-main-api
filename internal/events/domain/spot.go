package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSpotNameRequired    error = errors.New("spot name is required")
	ErrSpotNameLen         error = errors.New("spot name need to be only 2 characters")
	ErrSpotNameStart       error = errors.New("spot name must start with a letter")
	ErrSpotNameEnd         error = errors.New("spot name must end with a number")
	ErrSpotAlreadyreserved error = errors.New("spot already reserved")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketId string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyreserved
	}

	s.Status = SpotStatusSold
	s.TicketId = ticketID

	return nil
}

func (s *Spot) Validate() error {
	if s.Name == "" {
		return ErrSpotNameRequired
	}

	if len(s.Name) != 2 {
		return ErrSpotNameLen
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStart
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEnd
	}

	return nil
}

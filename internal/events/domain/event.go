package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired error = errors.New("event name is required")
	ErrEventDateFuture   error = errors.New("event date must be in the future")
	ErrEventCapacityZero error = errors.New("event capacity must be greater than zero")
	ErrEventPriceZero    error = errors.New("event price must be greather than zero")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Ticket       []Ticket
}

func (e *Event) Validate() error {
	if e.Name == "" {
		return errors.New("Event name is required")
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateFuture
	}

	if e.Capacity <= 0 {
		return ErrEventCapacityZero
	}

	if e.Price <= 0 {
		return ErrEventPriceZero
	}

	return nil
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)

	return spot, nil
}

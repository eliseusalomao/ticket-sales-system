package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired               = errors.New("event name is required")
	ErrEventDateIsBefore               = errors.New("event date must be in the future")
	ErrEventCapacitySize               = errors.New("event capacity must be greater than zero")
	ErrEventPriceMustBeGreaterThanZero = errors.New("event price must be greater than zero")
)

type Rating string

const (
	FreeRating Rating = "L"
	Rating10   Rating = "L10"
	Rating12   Rating = "L12"
	Rating14   Rating = "L14"
	Rating16   Rating = "L16"
	Rating18   Rating = "L18"
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
	Tickets      []Ticket
}

func (e *Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateIsBefore
	}

	if e.Capacity <= 0 {
		return ErrEventCapacitySize
	}

	if e.Price <= 0 {
		return ErrEventPriceMustBeGreaterThanZero
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

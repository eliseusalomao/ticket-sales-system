package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSpotNameIsRequired   = errors.New("spot name is required")
	ErrInvalidSpotNameLength       = errors.New("spot name must be at least 2 characters")
	ErrSpotNameMustStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotNameMustEndWithNumber   = errors.New("spot name must end with a number")
	ErrInvalidSpotNumber           = errors.New("invalid spot number")
	ErrSpotNotFound                = errors.New("spot not found")
	ErrSpotAlreadyTaken            = errors.New("spot already taken")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "Available"
	SpotStatusSold      SpotStatus = "Sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	validate := spot.Validate()
	if validate != nil {
		return nil, validate
	}

	return spot, nil
}

func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrInvalidSpotNameIsRequired
	}

	if len(s.Name) < 2 {
		return ErrInvalidSpotNameLength
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameMustStartWithLetter
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameMustEndWithNumber
	}

	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyTaken
	}
	s.Status = SpotStatusSold
	s.TicketID = ticketID
	return nil
}

package domain

import (
	"errors"
	"fmt"
)

type SpotService struct{}

var (
	ErrInvalidQuantity = errors.New("Invalid quantity. Must be greater than 0")
)

func NewSpotService() *SpotService {
	return &SpotService{}
}

func (s *SpotService) generateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}
	return nil
}

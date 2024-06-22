package repository

import (
	"database/sql"
	"errors"

	"github.com/ticket-system/golang/internal/events/domain"
)

type mysqlEventRepository struct {
	db *sql.DB
}

//func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
//	return &mysqlEventRepository{db:db}, nil
//}

func (r *mysqlEventRepository) CreateSpot(spot *domain.Spot) error {
	query := `
		INSERT INTO spots (id, event_id, name, status, ticket_id)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, spot.ID, spot.EventID, spot.Name, spot.Status, spot.TicketID)
	return err
}

func (r *mysqlEventRepository) ReserveSpot(spot *domain.Spot) error {
	query := `
		UPDATE spots
		SET status = ?, ticket_id = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query, domain.SpotStatusSold, spot.ID)
	return err
}

func (r *mysqlEventRepository) CreateTicket(ticket *domain.Ticket) error {
	query := `
		INSERT INTO tickets (id, event_id, spot_id, ticket_type, price)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, ticket.ID, ticket.EventID, ticket.Spot.ID, ticket.TicketType, ticket.Price)
	return err
}

func (r *mysqlEventRepository) FindEventById(eventID string) (*domain.Event, error) {
	query := `
		SELECT id, name, location, organization, rating, date, image_url, capacity, price, partner_id
		FROM events
		WHERE id = ?
	`

	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event *domain.Event
	err = rows.Scan(
		&event.ID,
		&event.Name,
		&event.Location,
		&event.Organization,
		&event.Rating,
		&event.Date,
		&event.ImageURL,
		&event.Capacity,
		&event.Price,
		&event.PartnerID,
		&event.Spots,
		&event.Tickets,
	)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *mysqlEventRepository) FindSpotsByEventID(eventID string) ([]*domain.Spot, error) {
	query := `
		SELECT id, event_id, name, status, ticket_id
		FROM spots
		WHERE event_id = ?
	`
	rows, err := r.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []*domain.Spot
	for rows.Next() {
		var spot domain.Spot
		if err := rows.Scan(
			&spot.ID, &spot.EventID, &spot.Name, &spot.Status, &spot.TicketID,
		); err != nil {
			return nil, err
		}
		spots = append(spots, &spot)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return spots, nil
}

func (r *mysqlEventRepository) FindSpotsByName(eventID, name string) (*domain.Spot, error) {
	query := `
		SELECT 
			s.id, s.event_id, s.name, s.status, s.ticket_id,
			t.id, t.event_id, t.spot_id, t.ticket_type, t.price
		FROM spots s
		LEFT JOIN tickets t ON s.id = t.spot_ideventID
		WHERE s.event_id = ? AND s.name = ?
	`
	row := r.db.QueryRow(query, eventID, name)

	var spot domain.Spot
	var ticket domain.Ticket
	var ticketID, ticketEventID, ticketSpotID, ticketType sql.NullString
	var ticketPrice sql.NullFloat64

	err := row.Scan(
		&spot.ID, &spot.EventID, &spot.Name, &spot.Status, &spot.TicketID,
		&ticketID, &ticketEventID, &ticketSpotID, &ticketType, &ticketPrice,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSpotNotFound
		}
		return nil, err
	}

	if ticketID.Valid {
		ticket.ID = ticketID.String
		ticket.EventID = ticketEventID.String
		ticket.Spot = &spot
		ticket.TicketType = domain.TicketType(ticketType.String)
		ticket.Price = ticketPrice.Float64
		spot.TicketID = ticket.ID
	}

	return &spot, nil
}

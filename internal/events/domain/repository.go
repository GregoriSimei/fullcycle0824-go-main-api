package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindById(eventID string) (*Event, error)
	FindSpotsByEventID(eventID string) ([]*Spot, error)
	FindASpotByName(eventID, spotName string) (*Spot, error)
	CreateEvent(event *Event) error
	CreateSpot(spot *Spot) error
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID, ticketID string) error
}

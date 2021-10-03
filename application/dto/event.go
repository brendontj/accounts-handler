package dto

type Event struct {
	EventType string `json:"type"`
	Destination *string `json:"destination,omitempty"`
	Origin *string `json:"origin,omitempty"`
	Amount int64 `json:"amount"`
}

type Response struct {
	Destination *AccountResponse `json:"destination,omitempty"`
	Origin *AccountResponse      `json:"origin,omitempty"`
}

type AccountResponse struct {
	ID string `json:"id"`
	Balance int64 `json:"balance"`
}
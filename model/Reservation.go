package model

type ReservationKey struct {
	ReservationID string `json:"reservationID"`
	RoomID        string `json:"roomID"`
}

type Reservation struct {
	ReservationKey
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

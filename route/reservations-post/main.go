package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/matheustex/go-reservation-api/model"
	"github.com/matheustex/go-reservation-api/service"
	"github.com/matheustex/go-reservation-api/util"
)

type ReservationRequest struct {
	RoomID    string `json:"roomID"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type Response struct {
	Reservation model.Reservation `json:"reservation"`
}

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	request := ReservationRequest{}
	err := json.Unmarshal([]byte(input.Body), &request)
	if err != nil {
		return util.NewErrorResponse(err)
	}

	if util.IsInvalidDate(request.StartDate, request.EndDate) {
		return util.NewErrorResponse(model.NewInputError("dates", "invalid"))
	}

	reservation := model.Reservation{
		ReservationKey: model.ReservationKey{
			RoomID:        request.RoomID,
			ReservationID: uuid.New().String(),
		},
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
	}

	err = service.PutReservation(&reservation)
	if err != nil {
		return util.NewErrorResponse(err)
	}

	return util.NewSuccessResponse(200, reservation)
}

func main() {
	lambda.Start(Handle)
}

package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matheustex/go-reservation-api/model"
	"github.com/matheustex/go-reservation-api/service"
	"github.com/matheustex/go-reservation-api/util"
)

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	roomID := input.PathParameters["id"]
	if len(roomID) == 0 {
		return util.NewErrorResponse(model.NewInputError("id", "invalid"))
	}

	reservationID := input.PathParameters["reservationId"]
	if len(reservationID) == 0 {
		return util.NewErrorResponse(model.NewInputError("reservationId", "invalid"))
	}

	err := service.DeleteReservation(roomID, reservationID)
	if err != nil {
		return util.NewErrorResponse(err)
	}

	return util.NewSuccessResponse(200, nil)
}

func main() {
	lambda.Start(Handle)
}

package main

import (
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matheustex/go-reservation-api/model"
	"github.com/matheustex/go-reservation-api/service"
	"github.com/matheustex/go-reservation-api/util"
)

func Handle(input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	offset, err := strconv.Atoi(input.QueryStringParameters["offset"])
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(input.QueryStringParameters["limit"])
	if err != nil {
		limit = 20
	}

	roomID := input.PathParameters["id"]
	if len(roomID) == 0 {
		return util.NewErrorResponse(model.NewInputError("id", "invalid"))
	}

	reservations, err := service.GetReservationsByRoomID(roomID, offset, limit)
	if err != nil {
		return util.NewErrorResponse(err)
	}

	return util.NewSuccessResponse(200, reservations)
}

func main() {
	lambda.Start(Handle)
}

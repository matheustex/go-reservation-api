package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/matheustex/go-reservation-api/model"
	"github.com/matheustex/go-reservation-api/util"
)

func PutReservation(reservation *model.Reservation) error {
	item, err := dynamodbattribute.MarshalMap(reservation)
	if err != nil {
		return err
	}

	nextReservations, err := GetNextReservationsByRoomID(reservation.RoomID, reservation.StartDate)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	if len(nextReservations) > 0 {
		refInterval := util.DateInterval{
			StartDate: reservation.StartDate,
			EndDate:   reservation.EndDate,
		}

		for _, item := range nextReservations {
			reservationInterval := util.DateInterval{
				StartDate: item.StartDate,
				EndDate:   item.EndDate,
			}

			if util.IsDateBooked(reservationInterval, refInterval) {
				err = model.NewInputError("Dates", "This date period is already booked.")
				return err
			}
		}
	}

	putReservation := dynamodb.PutItemInput{
		TableName: aws.String(ReservationTableName),
		Item:      item,
	}

	_, err = DynamoDB().PutItem(&putReservation)

	return err
}

func GetReservationsByRoomID(roomID string, offset, limit int) ([]model.Reservation, error) {
	queryReservationsByRoomID := dynamodb.QueryInput{
		TableName:                 aws.String(ReservationTableName),
		IndexName:                 aws.String("ByRoomID"),
		KeyConditionExpression:    aws.String("roomID=:roomID"),
		ExpressionAttributeValues: StringKey(":roomID", roomID),
		Limit:                     aws.Int64(int64(offset + limit)),
		ScanIndexForward:          aws.Bool(false),
	}

	items, err := QueryItems(&queryReservationsByRoomID, offset, limit)
	if err != nil {
		return nil, err
	}

	reservations := make([]model.Reservation, len(items))
	err = dynamodbattribute.UnmarshalListOfMaps(items, &reservations)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func GetNextReservationsByRoomID(roomID string, startDate string) ([]model.Reservation, error) {
	keyExpr := "roomID = :roomID AND endDate >= :startDate"
	keyMap := map[string]string{
		":roomID":    roomID,
		":startDate": startDate,
	}

	key, _ := dynamodbattribute.MarshalMap(keyMap)

	queryReservationsByRoomID := dynamodb.QueryInput{
		ConsistentRead:            aws.Bool(false),
		TableName:                 aws.String(ReservationTableName),
		IndexName:                 aws.String("ByRoomID"),
		KeyConditionExpression:    aws.String(keyExpr),
		ExpressionAttributeValues: key,
	}

	items, err := QueryAll(&queryReservationsByRoomID)
	if err != nil {
		return nil, err
	}

	reservations := make([]model.Reservation, len(items))
	err = dynamodbattribute.UnmarshalListOfMaps(items, &reservations)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func DeleteReservation(roomID string, reservationID string) error {
	key := model.ReservationKey{
		ReservationID: reservationID,
		RoomID:        roomID,
	}

	item, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		return err
	}

	deleteReservation := dynamodb.DeleteItemInput{
		TableName: aws.String(ReservationTableName),
		Key:       item,
	}

	_, err = DynamoDB().DeleteItem(&deleteReservation)

	return err
}

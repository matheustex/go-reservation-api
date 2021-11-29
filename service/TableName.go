package service

import (
	"fmt"
	"os"
)

var Stage = os.Getenv("STAGE")

var ReservationTableName = makeTableName("reservation")

func makeTableName(suffix string) string {
	return fmt.Sprintf("%s-%s", Stage, suffix)
}

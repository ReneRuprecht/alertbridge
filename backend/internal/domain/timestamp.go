package domain

import (
	"errors"
	"log"
	"time"
)

var ErrorTimestampInvalid = errors.New("Timestamp is invalid")

type Timestamp struct {
	time.Time
}

func NewTimestamp(startsAt string) (Timestamp, error) {

	s, err := time.Parse(time.RFC3339, startsAt)

	if err != nil {
		log.Println(ErrorTimestampInvalid.Error())
		return Timestamp{}, ErrorTimestampInvalid
	}
	return Timestamp{Time: s}, nil
}

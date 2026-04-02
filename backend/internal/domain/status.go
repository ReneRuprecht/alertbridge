package domain

import (
	"errors"
	"log"
)

var ErrorStatusEmpty = errors.New("Status cannot be emtpy")
var ErrorStatusInvalid = errors.New("Status is invalid")

type Status string

const (
	StatusFiring   Status = "firing"
	StatusResolved Status = "resolved"
)

func (s Status) isValid() bool {
	switch s {
	case StatusFiring, StatusResolved:
		return true
	default:
		return false
	}
}

func NewStatus(status string) (Status, error) {
	if status == "" {
		log.Println(ErrorStatusEmpty.Error())
		return "", ErrorStatusEmpty
	}

	s := Status(status)
	if !s.isValid() {
		log.Printf("Invalid status %s", status)
		return "", ErrorStatusInvalid
	}

	return s, nil

}

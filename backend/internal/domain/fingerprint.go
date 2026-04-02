package domain

import (
	"errors"
	"log"
)

var ErrorFingerprintEmpty = errors.New("Fingerprint cannot be emtpy")

type Fingerprint string

func NewFingerprint(fingerprint string) (Fingerprint, error) {
	if fingerprint == "" {
		log.Println(ErrorFingerprintEmpty.Error())
		return "", ErrorFingerprintEmpty
	}

	fp := Fingerprint(fingerprint)

	return fp, nil

}

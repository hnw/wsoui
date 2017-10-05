package wsoui

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

// Query the database for an entry based on the mac address
// If none are found ErrNotFound will be returned.
func LookUp(mac string) (string, error) {
	oui24, err := ParseMac(mac)
	if err != nil {
		return "", err
	}
	oui24hash := (int32(oui24[0])*256+int32(oui24[1]))*256 + int32(oui24[2])
	abbr, ok := oui[oui24hash]
	if !ok {
		return "", ErrNotFound
	}
	return abbr, nil
}

package wsoui

import (
	"bytes"
	"encoding/binary"
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
	oui24hash := (uint32(oui24[0])*256+uint32(oui24[1]))*256 + uint32(oui24[2])
	abbr64, ok := oui[oui24hash]
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, abbr64)
	abbr := string(bytes.TrimLeft(b, "\x00"))

	if !ok {
		return "", ErrNotFound
	}
	return abbr, nil
}

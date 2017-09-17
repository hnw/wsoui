package wsoui

import (
	"fmt"
	"strings"
)

// This error will be returned by ParseMac
// if the Mac address cannot be decoded.
type ErrInvalidMac struct {
	Reason string
	Mac    string
}

// Error returns a string representation of the error.
func (e ErrInvalidMac) Error() string {
	return "invalid mac address '" + e.Mac + "': " + e.Reason
}

// ParseMac will parse a string Mac address and return the first 3 entries.
// It will attempt to find a separator, ':' and '-' supported.
// If none of these are matched, it will assume there is none.
func ParseMac(mac string) (*[3]byte, error) {
	// Attempt to find a separator, ':' and '-' supported.
	if len(mac) < 6 {
		return nil, ErrInvalidMac{Reason: "Mac address too short. Should be at least 6 characters", Mac: mac}
	}
	var separator *byte
	var s []string

	if mac[2] == ':' || mac[2] == '-' {
		b := mac[2]
		separator = &b
		s = strings.Split(mac, string(*separator))
	} else {
		for i := 0; i < len(mac)-1; i += 2 {
			s = append(s, mac[i:i+2])
		}
	}
	if len(s) < 3 {
		return nil, ErrInvalidMac{Reason: "Unable to find at least 3 address elements", Mac: mac}
	}
	hw := [3]byte{}
	for i, p := range s {
		if i >= 3 {
			break
		}
		if len(p) != 2 {
			return nil, ErrInvalidMac{Reason: fmt.Sprintf("Address element %d (%s) is not 2 characters", i+1, p), Mac: mac}
		}
		var b byte
		n, err := fmt.Sscanf(p, "%x", &b)
		if n != 1 {
			return nil, ErrInvalidMac{Reason: fmt.Sprintf("Address element %d (%s) cannot be parsed as hex value: %v", i+1, p, err), Mac: mac}
		}
		hw[i] = b
	}
	return &hw, nil
}

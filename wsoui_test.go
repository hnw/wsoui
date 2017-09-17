package wsoui

import (
	"reflect"
	"testing"
)

func TestParseMac(t *testing.T) {
	expected := &[3]byte{0xcc, 0x20, 0xe8}
	oui24, err := ParseMac("cc-20-e8")
	if err != nil {
		t.Errorf("ParseMac() Failure: %s", err)
	} else if !reflect.DeepEqual(expected, oui24) {
		t.Errorf("Expected %#v, got %#v", expected, oui24)
	}
}

func TestLookUp(t *testing.T) {
	expected := "Apple"
	abbr, err := LookUp("cc-20-e8")
	if err != nil {
		t.Errorf("LookUp() Failure: %s", err)
	} else if !reflect.DeepEqual(expected, abbr) {
		t.Errorf("Expected %#v, got %#v", expected, abbr)
	}
}

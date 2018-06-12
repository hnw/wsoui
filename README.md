wsoui: The OUI database for Go [![Build Status](https://travis-ci.org/hnw/wsoui.svg?branch=master)](https://travis-ci.org/hnw/wsoui)
=====================

## Description

`wsoui` is a Go library which provides the vendor name lookup from MAC address.

This file was auto-generated from [Wireshark manufacturer database](https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf), so you should follow original license.

## Usage

### `func LookUp(mac string) (string, error)`

This function returns the vendor abbreviation name (like `"Cisco"` ) corresponds to the argument `mac` which is the first three octets of MAC address.


## Sample

``` go
package main

import (
	"fmt"
	"github.com/hnw/wsoui"
)

func main() {
	abbr, _ := wsoui.LookUp("cc-20-e8")
	fmt.Print(abbr) // Apple
}
```

package dev

import (
	"github.com/ge-lighting/ble"
	"github.com/ge-lighting/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}

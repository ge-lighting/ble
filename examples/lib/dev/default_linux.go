package dev

import (
	"github.com/ge-lighting/ble"
	"github.com/ge-lighting/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}

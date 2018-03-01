// +build !linux,!darwin,!windows,!openbsd,!freebsd

package water

import "errors"

func newTAP(config Config) (ifce *Interface, err error) {
	return nil, errors.New("tap interface not implemented on this platform")
}

func newTUN(config Config) (ifce *Interface, err error) {
	return nil, errors.New("tap interface not implemented on this platform")
}

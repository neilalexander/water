// +build openbsd freebsd

package water

import (
	"os"
)

const bsdTUNSIFINFO = (0x80000000) | ((4 & 0x1fff) << 16) | uint32(byte('t'))<<8 | 91
const bsdTUNSIFMODE = (0x80000000) | ((4 & 0x1fff) << 16) | uint32(byte('t'))<<8 | 94

type tuninfo struct {
	BaudRate int
	MTU      int16
	Type     uint8
	Dummy    uint8
}

func newTAP(config Config) (ifce *Interface, err error) {
	if config.Name[:8] != "/dev/tap" {
		panic("TUN/TAP name must be in format /dev/tunX or /dev/tapX")
	}

	file, err := os.OpenFile(config.Name, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	ifce = &Interface{isTAP: true, ReadWriteCloser: file, name: config.Name[5:]}
	return
}

func newTUN(config Config) (ifce *Interface, err error) {
	if config.Name[:8] != "/dev/tun" {
		panic("TUN/TAP name must be in format /dev/tunX or /dev/tapX")
	}

	file, err := os.OpenFile(config.Name, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	ifce = &Interface{isTAP: false, ReadWriteCloser: file, name: config.Name[5:]}
	return
}

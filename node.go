package uuid

import (
	"bytes"
	"net"
)

func getNode() string {
	node, ok := getMacAddr()
	if !ok {
		return "00:00:00:00:00:00"
	}

	return node
}

func getMacAddr() (string, bool) {
	var addr string
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				addr = i.HardwareAddr.String()
				return addr, true
			}
		}
	}
	return addr, false
}

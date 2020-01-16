package uuid

import (
	"bytes"
	"net"
)

var node [6]byte

func init() {
	node = getNode()
}

func getNode() [6]byte {
	var nodeID [6]byte

	node, ok := getMacAddr()
	if !ok {
		return nodeID
	}

	copy(nodeID[:], node[:6])
	return nodeID
}

func getMacAddr() ([]byte, bool) {
	var addr []byte
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				addr = i.HardwareAddr
				return addr, true
			}
		}
	}
	return addr, false
}

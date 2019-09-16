package uuid

import (
	"encoding/binary"
	"encoding/hex"
)

// NewUUID create Universally unique identifier
func NewUUID() string {
	var uuid [16]byte
	t := getTimeSince1582()
	cSeq := clockSeq()
	timeLow := uint32(t)
	timeMid := uint16((t >> 32))
	timeHi := uint16((t >> 48))
	timeHi += 0x1000

	node := getNode()

	binary.BigEndian.PutUint32(uuid[0:], timeLow)
	binary.BigEndian.PutUint16(uuid[4:], timeMid)
	binary.BigEndian.PutUint16(uuid[6:], timeHi)
	binary.BigEndian.PutUint16(uuid[6:], timeHi)
	binary.BigEndian.PutUint16(uuid[8:], cSeq)

	copy(uuid[10:], node[:6])

	return encode(uuid)
}

func encode(uuid [16]byte) string {
	dst := make([]byte, hex.EncodedLen(len(uuid)+3))

	hex.Encode(dst, uuid[0:4])
	dst[8] = '-'
	hex.Encode(dst[9:17], uuid[4:8])
	dst[17] = '-'
	hex.Encode(dst[18:26], uuid[8:12])
	dst[26] = '-'
	hex.Encode(dst[27:], uuid[12:])

	return string(dst[:])
}

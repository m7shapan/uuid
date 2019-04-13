package uuid

import (
	"fmt"
	"strconv"
	"strings"
)

// NewUUID create Universally unique identifier
func NewUUID() string {

	binaryTime := get60BitBinaryTime()
	clockSeq := get14BitClockSeq()
	node := getNode()
	firstNodeID, secondNodeID := getNodeID(node)

	return fmt.Sprintf("%s-%s%s-%s%s%s-%s",
		getTimeLow(binaryTime),
		getTimeMid(binaryTime),
		getTimeHiAndVersion(binaryTime),
		getClockSeqHiAndReserved(clockSeq),
		getClockSeqLow(clockSeq),
		firstNodeID,
		secondNodeID,
	)
}

func getTimeLow(s string) string {
	i, _ := strconv.ParseUint(s[0:32], 2, 32)
	return fmt.Sprintf("%x", i)
}

func getTimeMid(s string) string {
	i, _ := strconv.ParseUint(s[32:48], 2, 32)
	return fmt.Sprintf("%x", i)
}

func getTimeHiAndVersion(s string) string {
	i, _ := strconv.ParseUint(fmt.Sprintf("%s%s", "00011", s[48:60]), 2, 32)
	return fmt.Sprintf("%x", i)
}

func getClockSeqLow(s string) string {
	i, _ := strconv.ParseUint(s[0:8], 2, 8)
	return fmt.Sprintf("%x", i)
}

func getClockSeqHiAndReserved(s string) string {
	i, _ := strconv.ParseUint(fmt.Sprintf("%s%s", "10", s[8:14]), 2, 8)
	return fmt.Sprintf("%x", i)
}

func getNodeID(s string) (string, string) {
	sArr := strings.Split(s, ":")

	return fmt.Sprintf("%s", strings.Join(sArr[0:2], "")), fmt.Sprintf("%s", strings.Join(sArr[2:], ""))
}

package uuid

import (
	"math/rand"
	"sync"
	"time"
)

var lastTime uint64
var clockSeq uint16

var lock = sync.RWMutex{}

func init() {
	// 16383 is the max number of 14 bit
	clockSeq = uint16(rand.Intn(16383))
}

func time100Nano() uint64 {
	t := time.Now()
	return uint64(t.UnixNano() / 100)
}

func getTimeSince1582() uint64 {

	return time100Nano() + 122192928000000000
}

func getTime() (uint64, uint16) {
	defer lock.Unlock()
	lock.Lock()

	t := getTimeSince1582()
	if lastTime == t {
		clockSeq++
	} else {
		lastTime = t
		clockSeq = 0
	}

	return t, clockSeq
}

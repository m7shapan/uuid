package uuid

import (
	"sync"
	"time"
)

var lastTime uint64
var clockSeq uint16

var lock = sync.RWMutex{}

func getNanosecond() int64 {
	defer lock.Unlock()

	lock.Lock()
	t := time.Now()
	return t.UnixNano()
}

func time100Nano() uint64 {
	return uint64(getNanosecond() / 100)
}

func getTimeSince1582() uint64 {

	return time100Nano() + 122192928000000000
}

func getTime() (uint64, uint16) {
	t := getTimeSince1582()

	defer lock.Unlock()
	lock.Lock()
	if lastTime == t {
		clockSeq++
	} else {
		lastTime = t
		clockSeq = 0
	}

	return t, clockSeq
}

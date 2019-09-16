package uuid

import (
	"math/rand"
	"sync"
	"time"
)

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

func clockSeq() uint16 {
	// 16383 is the max number of 14 bit
	return uint16(rand.Intn(16383))
}

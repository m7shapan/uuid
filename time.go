package uuid

import (
	"math/rand"
	"strconv"
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

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func get60BitBinaryTime() (s string) {
	s = strconv.FormatUint(getTimeSince1582(), 2)

	for index := len(s); index < 60; index++ {
		s = "0" + s
	}

	s = reverse(s)
	return
}

func clockSeq() uint64 {
	// 16383 is the max number of 14 bit
	return uint64(rand.Intn(16383))
}

func get14BitClockSeq() (s string) {
	s = strconv.FormatUint(clockSeq(), 2)

	for index := len(s); index < 14; index++ {
		s = "0" + s
	}

	s = reverse(s)
	return
}

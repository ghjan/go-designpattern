package counter

import "sync"

type Counter struct {
	sync.Mutex
	Value int
}

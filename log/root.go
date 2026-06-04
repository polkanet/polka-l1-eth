package log

import (
	"log/slog"
	"os"
	"sync"
)

var (
	rootLock sync.RWMutex
	root 	 Logger
)

func initializeRoot() {
	root = &logger{
		slog.New(DiscardHandler()),
	}
}

func SetDefault(l Logger) {
	rootLock.Lock()
	defer rootLock.Unlock()

	root = l
}
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


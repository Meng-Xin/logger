package hooks

import (
	"go.uber.org/zap/zapcore"
	"sync/atomic"
)

func makeCountingHook() (func(zapcore.Entry) error, *atomic.Int64) {
	count := &atomic.Int64{}
	h := func(zapcore.Entry) error {
		count.Add(1)
		return nil
	}
	return h, count
}

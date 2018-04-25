package util

import (
	"github.com/Sharykhin/it-customer-review/tone-analyzer/logger"
)

// Check is a helper func that checks deferred call on errors
func Check(fn func() error) {
	if err := fn(); err != nil {
		logger.Logger.Errorf("deferred calls returned an error:%v", err)
	}
}

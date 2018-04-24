package logger

import (
	"github.com/Sharykhin/it-customer-review/api/logger/file"
	"github.com/Sharykhin/it-customer-review/tone-analyzer/contract"
)

type (
	logger struct {
		contract.Logger
	}
)

// Logger is a general service that returns concrete one which implement logger interface
var Logger = logger{Logger: file.Logger}

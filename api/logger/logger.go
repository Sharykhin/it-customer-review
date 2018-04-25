package logger

import (
	"github.com/Sharykhin/it-customer-review/api/contract"
	"github.com/Sharykhin/it-customer-review/api/logger/file"
)

type (
	logger struct {
		contract.Logger
	}
)

// Logger is a general service that returns concrete one which implement logger interface
var Logger = logger{Logger: file.Logger}

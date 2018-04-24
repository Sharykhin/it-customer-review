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

var Logger = logger{Logger: file.FileLogger}

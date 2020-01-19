package date_utils

import (
	"time"
)

const (
	apiStringLayout   = "2006-01-02T15:04:05Z"
	apiDatabaseLayout = "2006-01-02 15:04:05"
)

func CurrentDate() time.Time {
	return time.Now().UTC()
}

func CurrentDateDefaultFormat() string {
	return CurrentDate().Format(apiStringLayout)
}

func CurrentDateDbFormat() string {
	return CurrentDate().Format(apiDatabaseLayout)
}

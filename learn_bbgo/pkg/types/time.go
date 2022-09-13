package types

import (
	"time"
)

// Time type implements the driver value for sqlite
type Time time.Time

func (t Time) Time() time.Time {
	return time.Time(t)
}

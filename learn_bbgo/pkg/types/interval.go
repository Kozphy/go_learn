package types

type Interval string

// IntervalWindow is used by the indicators
type IntervalWindow struct {
	// The interval of kline
	Interval Interval `json:"interval"`

	// The windows size of the indicator (EWMA and SMA)
	Window int `json:"window"`
}

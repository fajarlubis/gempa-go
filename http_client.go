package gempago

import (
	"errors"
	"time"
)

type WithTimeout struct {
	max, min  time.Duration
	timeout   time.Duration
	factor    float64
	startedAt time.Time
	nextDelay float64
}

func NewRequestWithTimeout(max, min time.Duration, timeout time.Duration, factor float64, startedAt time.Time) (*WithTimeout, error) {
	if factor >= 1.0 || factor <= 0.0 {
		return nil, errors.New("factor should be between 0 and 1")
	}

	ieb := WithTimeout{
		max:       max,
		min:       min,
		timeout:   timeout,
		factor:    factor,
		startedAt: startedAt,
		nextDelay: nextDelay,
	}

	return &ieb, nil
}

func (m *WithTimeout) Next() error {

}

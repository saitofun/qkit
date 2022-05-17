package types

import (
	"strconv"
	"time"
)

type Second Duration

var _ Span = Second(0)

func (s *Second) UnmarshalText(text []byte) error {
	d := Duration(0)
	if err := d.UnmarshalText(text); err != nil {
		return err
	}
	*s = Second((time.Duration)(d).Seconds())
	return nil
}

func (s Second) MarshalText() ([]byte, error) {
	return []byte(s.Literal()), nil
}

func (s Second) Duration() time.Duration { return time.Duration(s) * time.Second }

func (s Second) Int() int64 { return int64(s) }

func (s Second) String() string { return s.Duration().String() }

// Literal is present second literal string, `1 second` presents as `1`
func (s Second) Literal() string { return strconv.FormatInt(int64(s), 10) }

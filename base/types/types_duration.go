package types

import "time"

type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}
	du, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = AsDuration(du)
	return nil
}

func (d Duration) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d Duration) Duration() time.Duration { return time.Duration(d) }

func (d Duration) Int() int64 { return int64(d) }

func (d Duration) String() string { return d.Duration().String() }

func AsDuration(d time.Duration) Duration { return Duration(d) }

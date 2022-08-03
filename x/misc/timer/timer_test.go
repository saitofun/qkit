package timer_test

import (
	"testing"
	"time"

	"github.com/saitofun/qkit/x/misc/timer"
)

func TestStartSpan(t *testing.T) {
	span := timer.StartSpan()

	t.Log(span.StartedAt())
	time.Sleep(time.Second)
	t.Log(span.Cost())
	time.Sleep(time.Second)
	t.Log(span.Cost())
}

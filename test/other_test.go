package test

import (
	"testing"
	"time"
)

func TestUnixNano(t *testing.T) {
	now := time.Now()

	t.Log(now.UnixNano())

	// convert unix nano to time
	me := time.Unix(0, now.UnixNano())
	time.Sleep(5 * time.Second)
	t.Log(me.UnixNano())

	now = time.Now()
	t.Log(now.UnixNano())

	t.Log(now.Sub(me) >= 5*time.Second)
}

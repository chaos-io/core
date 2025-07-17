package core

import (
	"math"
	"time"
)

func FromDuration(d time.Duration) *Duration {
	dur := &Duration{}
	return dur.FromDuration(d)
}

func NewDuration(sec float64) *Duration {
	dur := &Duration{}
	return dur.FromSeconds(sec)
}

func (x *Duration) FromDuration(d time.Duration) *Duration {
	if x != nil {
		x.Seconds = int64(d.Seconds())
		x.Nanoseconds = int32(int64(d) - x.Seconds*int64(time.Second))
	}
	return x
}

func (x *Duration) FromSeconds(sec float64) *Duration {
	if x != nil {
		x.Seconds = int64(sec)
		delta := sec - float64(x.Seconds)
		x.Nanoseconds = int32(math.Round(delta * float64(time.Second)))
	}
	return x
}

package core

import "time"

func Now() *Timestamp {
	return FromTime(time.Now())
}

// FromTime covert form time.Time to Timestamp
func FromTime(t time.Time) *Timestamp {
	sec := t.Unix()
	return &Timestamp{
		Seconds:     sec,
		Nanoseconds: int32(t.UnixNano() - sec*1e9),
	}
}

func Since(t *Timestamp) *Duration {
	if t != nil {
		return FromDuration(time.Since(t.ToTime()))
	}
	return nil
}

func Until(t *Timestamp) *Duration {
	if t != nil {
		return FromDuration(time.Until(t.ToTime()))
	}
	return nil
}

func (x *Timestamp) FromTime(t time.Time) *Timestamp {
	if x != nil {
		ft := FromTime(t)
		x.Seconds = ft.Seconds
		x.Nanoseconds = ft.Nanoseconds
	}
	return x
}

func (x *Timestamp) ToTime() time.Time {
	if x != nil {
		return time.Unix(x.Seconds, int64(x.Nanoseconds))
	}
	return time.Time{}
}

func (x *Timestamp) After(u *Timestamp) bool {
	if x != nil && u != nil {
		if x != u {
			return x.ToTime().After(u.ToTime())
		}
	}
	return false
}

func (x *Timestamp) Before(u *Timestamp) bool {
	if x != nil && u != nil {
		if x != u {
			return x.ToTime().Before(u.ToTime())
		}
	}
	return false
}

func (x *Timestamp) Equal(u *Timestamp) bool {
	if x != nil && u != nil {
		if x == u {
			return true
		}
		return x.ToTime().Equal(u.ToTime())
	}
	return false
}

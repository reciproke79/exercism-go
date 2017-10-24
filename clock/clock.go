package clock

import (
	"time"
)

const testVersion = 4

type Clock struct {
	Time time.Time
}

func New(hour, minute int) Clock {
	// Define a template for hours and minutes
	t := time.Date(0, 0, 0, hour, minute, 0, 0, time.Local)
	return Clock{
		Time: time.Date(2009, 1, 1, t.Hour(), t.Minute(), 0, 0, time.Local),
	}
}

func (c Clock) String() string {
	return c.Time.Format("15:04")
}

func (c Clock) Add(minutes int) Clock {
	return Clock{
		Time: c.Time.Add(time.Minute * time.Duration(minutes)),
	}
}

package hyperscript

import "strconv"

type TimeUnit string

const (
	TimeUnitSeconds      TimeUnit = "s"
	TimeUnitMilliseconds TimeUnit = "ms"
)

type Time struct {
	number int
	unit   TimeUnit
}

func (t *Time) String() string {
	return strconv.Itoa(t.number) + string(t.unit)
}

func TimeSeconds(s int) *Time {
	return &Time{number: s, unit: TimeUnitSeconds}
}

func TimeMilliseconds(ms int) *Time {
	return &Time{number: ms, unit: TimeUnitMilliseconds}
}

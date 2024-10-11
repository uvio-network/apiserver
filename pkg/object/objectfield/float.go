package objectfield

import "time"

type Float struct {
	// Data is the float data of this object field.
	Data float64 `json:"data,omitempty"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time,omitempty"`
}

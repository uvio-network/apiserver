package objectfield

import "time"

type ListFloat struct {
	// Data is the list data of this object field.
	Data []float64 `json:"data,omitempty"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time,omitempty"`
}

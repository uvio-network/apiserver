package objectfield

import "time"

type MapFloat struct {
	// Data is the map data of this object field.
	Data map[string]float64 `json:"data,omitempty"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time,omitempty"`
}

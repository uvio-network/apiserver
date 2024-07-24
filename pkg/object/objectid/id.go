package objectid

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/xh3b4sd/tracer"
)

type ID string

func Random(num int64) ID {
	return ID(fmt.Sprintf("%d%06d", num, rand.Intn(999999)))
}

func (i ID) Float() float64 {
	f, e := strconv.ParseFloat(string(i), 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return f
}

func (i ID) Int() int64 {
	n, e := strconv.ParseInt(string(i), 10, 64)
	if e != nil {
		tracer.Panic(tracer.Mask(e))
	}

	return n
}

func (i ID) String() string {
	return string(i)
}

// Verify returns true if the underlying ID is greater than a random ID that has
// been generated in the past. This verification check works as long as the
// underlying ID is derived by the current time of ID generation. Verify ensures
// that an ID cannot simply be empty, while following our currently implemented
// format.
func (i ID) Verify() bool {
	return i.Int() > 1721826156886308
}

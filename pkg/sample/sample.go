package sample

import (
	"fmt"
	"io"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Rea io.Reader
}

type Sample struct {
	rea io.Reader
}

func New(c Config) *Sample {
	if c.Rea == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rea must not be empty", c)))
	}

	var s *Sample
	{
		s = &Sample{
			rea: c.Rea,
		}
	}

	return s
}

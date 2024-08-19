package sample

type Interface interface {
	// Random returns a randomized sample of the points on the lines defined by
	// the input lengths. Randomization happens according to the underlying source
	// of entropy. The amount of returned points must be equal on either side and
	// may be capped at some maximum.
	//
	//     inp[0] the maximum length of line A
	//     inp[1] the maximum length of line B
	//     out[2] the random points on line A
	//     out[3] the random points on line B
	//
	Random(uint, uint) ([]int, []int)
}

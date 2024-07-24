package poststorage

type Interface interface {
	// Create persists new post objects.
	//
	//     @inp[0] the post objects to create
	//     @out[0] the post objects mapped to their internal post IDs
	//
	Create([]*Object) ([]*Object, error)
}

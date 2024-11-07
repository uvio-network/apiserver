package notestorage

import "github.com/xh3b4sd/objectid"

type Interface interface {
	// CreateNote persists new note objects in the underlying storage.
	//
	//     @inp[0] the note objects to create
	//
	CreateNote([]*Object) error

	// SearchPage returns the note objects within the given pagination range in
	// reversed order of claim creation time. Given the indexed notes [A B C D E],
	// the first page [0 3] returns the most recent notes [C D E].
	//
	//     @inp[0] the user ID to search notes for
	//     @inp[1] the list of note kinds to search notes for
	//     @inp[2] the start paging pointer defining the beginning of the page
	//     @inp[3] the stop paging pointer defining the end of the page
	//     @out[0] the list of note objects within the given pagination range
	//
	SearchPage(objectid.ID, []string, int, int) ([]*Object, error)

	// SearchTime returns the note objects within the given pagination range of
	// creation time. Given the indexed notes [A B C D E], the time range
	// [1730761200 1730847600] returns the respective notes [B C D].
	//
	//     @inp[0] the user ID to search notes for
	//     @inp[1] the list of note kinds to search notes for
	//     @inp[2] the start unix timestamp defining the beginning of the page
	//     @inp[3] the stop unix timestamp defining the end of the page
	//     @out[0] the list of note objects within the given pagination range
	//
	SearchTime(objectid.ID, []string, int64, int64) ([]*Object, error)
}

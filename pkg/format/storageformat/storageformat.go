package storageformat

const (
	// PostObject is used to store our internal representation of a post object.
	//
	//     post ID                   post object
	//                         ->
	//     post/object/1234          {"key": "val"}
	//
	PostObject = "post/object/%s"
)

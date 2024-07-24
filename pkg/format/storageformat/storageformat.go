package storageformat

const (
	// PostObject is used to store our internal representation of a post object.
	//
	//     post ID                   post object
	//                         ->
	//     post/object/1234          {"key": "val"}
	//
	PostObject = "post/object/%s"

	// PostTree is used to store all claim IDs that belong to the same tree.
	//
	//     tree ID                 claim IDs
	//                       ->
	//     post/tree/1234          3456,5678
	//
	PostTree = "post/tree/%s"
)

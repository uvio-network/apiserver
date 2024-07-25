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

	// UserSubject is used to store the user specific mappings between external
	// and internal identity representations. An external representation might be
	// an OAuth subject claim provided with an access token when authenticating
	// via Google. This subject claim would become part of the key used here. The
	// internal user representation is our own unified user ID, which would then
	// become the value stored using the created subject claim key.
	//
	//     external subject claim                   internal user ID
	//                                        ->
	//     user/subject/google-oauth2|1234          5678
	//
	UserSubject = "user/subject/%s"

	// UserObject is used to store our internal representation of a user object.
	//
	//     user ID                   user object
	//                         ->
	//     user/object/1234          {"key": "val"}
	//
	UserObject = "user/object/%s"
)

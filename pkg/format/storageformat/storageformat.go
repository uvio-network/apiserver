package storageformat

const (
	// PostCreated is used to store all claim IDs as they have been created in
	// chronological order.
	//
	//     symbol key            claim IDs
	//                     ->
	//     post/created          3456,5678
	//
	PostCreated = "post/created"

	// PostLabel is used to store all claim IDs associated to a given label name.
	// This key allows us to search for all claims that are categorized under any
	// given label.
	//
	//     label name              claim IDs
	//                       ->
	//     post/label/mev          3456,5678
	//
	PostLabel = "post/label/%s"

	// PostObject is used to store our internal representation of a post object.
	// This key allows us to search for all claim objects by their own object ID.
	//
	//     post ID                   post object
	//                         ->
	//     post/object/1234          {"key": "val"}
	//
	PostObject = "post/object/%s"

	// PostTree is used to store all claim IDs that belong to the same tree. This
	// key allows us to search for all claims that belong to the same claim tree.
	//
	//     tree ID                 claim IDs
	//                       ->
	//     post/tree/1234          3456,5678
	//
	PostTree = "post/tree/%s"

	// PostOwner is used to store all claim IDs that have been created by a
	// specific user. This key allows us to search for all claims that have been
	// created by any given user.
	//
	//     user ID                  claim IDs
	//                        ->
	//     post/owner/1234          3456,5678
	//
	PostOwner = "post/owner/%s"

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
	// This key allows us to search for all user objects by their own object ID.
	//
	//     user ID                   user object
	//                         ->
	//     user/object/1234          {"key": "val"}
	//
	UserObject = "user/object/%s"

	// VoteObject is used to store our internal representation of a vote object.
	// This key allows us to search for all vote objects by their own object ID.
	//
	//     vote ID                   vote object
	//                         ->
	//     vote/object/1234          {"key": "val"}
	//
	VoteObject = "vote/object/%s"

	// VoteOwner is used to store all vote IDs that have been created by a
	// specific user. This key allows us to search for all votes that have been
	// created by any given user.
	//
	//     user ID                  vote IDs
	//                        ->
	//     vote/owner/1234          3456,5678
	//
	VoteOwner = "vote/owner/%s"
)

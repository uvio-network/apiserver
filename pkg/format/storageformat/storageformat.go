package storageformat

const (
	// NoteObject is used to store our internal representation of a note object.
	// This key allows us to search for all note objects by their own object ID.
	//
	//     note ID                   note object
	//                         ->
	//     note/object/1234          {"key": "val"}
	//
	NoteObject = "note/object/%s"

	// NoteOwner is used to store all note IDs that have been created for a
	// specific user. This key allows us to search for all notes that belong to
	// any given user.
	//
	//     user ID                  note IDs
	//                        ->
	//     note/owner/1234          3456,5678
	//
	NoteOwner = "note/owner/%s"

	// PostComment is used to store all comment IDs posted on a specific parent
	// claim. This key allows us to search for all comments on any given market.
	//
	//     claim ID                               comment IDs
	//                                      ->
	//     post/kind/comment/parent/1234          3456,5678
	//
	PostComment = "post/kind/comment/parent/%s"

	// PostExpiry is used to store all claim IDs by claim expiry as defined in the
	// created claim objects. This key helps us to automatically progress in claim
	// trees, for creating claims of kind "resolve" once a claim of kind "propose"
	// expires.
	//
	//     symbol key                             claim IDs
	//                                      ->
	//     post/lifecycle/propose/expiry          3456,5678
	//
	PostExpiry = "post/lifecycle/%s/expiry"

	// PostLabel is used to store all claim IDs associated to a given label name.
	// This key allows us to search for all claims that are categorized under any
	// given label.
	//
	//     label name              claim IDs
	//                       ->
	//     post/label/mev          3456,5678
	//
	PostLabel = "post/label/%s"

	// PostLifecycle is used to store all claim IDs by their respective lifecycle
	// configuration. Using that mapping we can search e.g. for all disputes on
	// the platform.
	//
	//     lifecycle name                  claim IDs
	//                               ->
	//     post/lifecycle/propose          3456,5678
	//
	PostLifecycle = "post/lifecycle/%s"

	// PostObject is used to store our internal representation of a post object.
	// This key allows us to search for all claim objects by their own object ID.
	//
	//     post ID                   post object
	//                         ->
	//     post/object/1234          {"key": "val"}
	//
	PostObject = "post/object/%s"

	// PostOwner is used to store all claim IDs that have been created by a
	// specific user. This key allows us to search for all claims that have been
	// created by any given user.
	//
	//     user ID                  claim IDs
	//                        ->
	//     post/owner/1234          3456,5678
	//
	PostOwner = "post/owner/%s"

	// PostTree is used to store all claim IDs that belong to the same tree. This
	// key allows us to search for all claims that belong to the same claim tree.
	//
	//     tree ID                 claim IDs
	//                       ->
	//     post/tree/1234          3456,5678
	//
	PostTree = "post/tree/%s"

	// PostOwnerComment is used to store all comment IDs created by a specific
	// user and commented on a specific claim. This key allows us to search for
	// all comments a user has made on any given market that they participated in.
	//
	//     user ID / claim ID                                comment IDs
	//                                                 ->
	//     post/kind/comment/owner/1234/parent/1234          3456,5678
	//
	PostOwnerComment = "post/kind/comment/owner/%s/parent/%s"

	// UserReputation is used to store the user IDs with the highest reputation
	// ranks on the platform. The amount of users tracked here may be limited.
	//
	//     symbol key               user IDs
	//                        ->
	//     user/reputation          3456,5678
	//
	UserReputation = "user/reputation"

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

	// VoteClaim is used to store all vote IDs for votes that have been cast on a
	// specific claim. This key allows us to search for all votes associated to
	// any given claim.
	//
	//     claim ID                 vote IDs
	//                        ->
	//     vote/claim/1234          3456,5678
	//
	VoteClaim = "vote/claim/%s"

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

	// VoteOwnerClaim is used to store all vote IDs that have been created by a
	// specific user on a specific claim. This key allows us to search for all
	// votes any given user made on any given claim.
	//
	//     user ID / claim ID                  vote IDs
	//                                   ->
	//     vote/owner/1234/claim/1234          3456,5678
	//
	VoteOwnerClaim = "vote/owner/%s/claim/%s"

	// WalletAddress is used to store all wallet IDs that are associated to a
	// particular wallet address. This key allows us to search for all user IDs
	// given any wallet address.
	//
	//     wallet address                 wallet IDs
	//                              ->
	//     wallet/address/0x1234          3456,5678
	//
	WalletAddress = "wallet/address/%s"

	// WalletObject is used to store our internal representation of a wallet
	// object. This key allows us to search for all wallet objects by their own
	// object ID.
	//
	//     wallet ID                   wallet object
	//                           ->
	//     wallet/object/1234          {"key": "val"}
	//
	WalletObject = "wallet/object/%s"

	// WalletOwner is used to store all wallet IDs that have been created by a
	// specific user. This key allows us to search for all wallets that have been
	// created by any given user.
	//
	//     user ID                    wallet IDs
	//                          ->
	//     wallet/owner/1234          3456,5678
	//
	WalletOwner = "wallet/owner/%s"
)

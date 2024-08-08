package poststorage

import (
	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreatePost(inp []*Object) error {
	var err error

	for i := range inp {
		// Create the normalized key-value pair for the post object. With that we
		// can search for post objects using their IDs.
		{
			err = r.red.Simple().Create().Element(posObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Store every claim ID in a global sorted set, based on their time of
		// creation. This step ensures we can search for claims using pagination in
		// reverse chronolical order.
		{
			err = r.red.Sorted().Create().Score(storageformat.PostCreated, inp[i].ID.String(), float64(inp[i].Created.UnixNano()))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Add the given claim object ID to the list of label names, so that we can
		// search for claims given any label name.
		for _, x := range inp[i].Labels {
			err = r.red.Sorted().Create().Score(posLab(x), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Associate every given post ID with the given tree ID. That means we are
		// able to search for all posts of a given tree at once, including claims
		// and comments.
		{
			err = r.red.Sorted().Create().Score(posTre(inp[i].Tree), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Track the user creating this post as the owner, and make sure that we can
		// find all posts for any given user ID.
		{
			err = r.red.Sorted().Create().Score(posOwn(inp[i].Owner), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if inp[i].Kind == "comment" {
			// Store the given comment ID under a key that links to its parent claim.
			// Storing this relationship enables us to search for all comments that
			// all users made on a specific market.
			err = r.red.Sorted().Create().Score(posCom(inp[i].Parent), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}

			// Store the given post ID under a key that combines the post owner and
			// the post parent, if the given post is in fact a comment. Storing this
			// relationship enables us to search for comments that users made on
			// markets in which they have skin in the game.
			err = r.red.Sorted().Create().Score(posUseCom(inp[i].Owner, inp[i].Parent), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

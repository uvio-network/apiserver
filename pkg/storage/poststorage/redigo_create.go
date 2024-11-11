package poststorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/tracer"
)

const (
	oneWeek = time.Hour * 24 * 7
)

func (r *Redigo) CreateExpiry(inp []*Object) error {
	var err error

	for i := range inp {
		if inp[i].Kind == "claim" {
			err = r.red.Sorted().Create().Score(posExp(inp[i].Lifecycle.Data), inp[i].ID.String(), float64(inp[i].Expiry.Unix()))
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func (r *Redigo) CreatePost(inp []*Object) error {
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

		// Keep track of all claim expiries, except for claims with lifecycle phase
		// "balance", because balances are final and do not have an expiry.
		if inp[i].Kind == "claim" && !inp[i].Lifecycle.Is(objectlabel.LifecycleBalance) {
			var exp time.Time
			{
				exp = inp[i].Expiry
			}

			// Claims of lifecycle phase "resolve" have two stages of expiries. The
			// first stage is the defined expiry that we track in the post object.
			// Within this first stage it is possible for voters to cast their votes.
			// Once the first stage concluded, the second stage begins, which is the
			// challenge window during which disputes may be created by anyone. When
			// we expire resolves internally, then we aim to update balances in order
			// to settle a claim tree. So the resolve expiry tracked in our internal
			// expiry queue considers the claim's challenge window, which is hard
			// coded system wide to be 7 standard days. If no dispute has been
			// proposed within this challenge window, then we can settle the original
			// claim using our background automation.
			if inp[i].Lifecycle.Is(objectlabel.LifecycleResolve) {
				exp = exp.Add(oneWeek)
			}

			// Regardless the claim lifecycle, we keep track of all claim IDs by their
			// specified claim expiry. This helps us to automate the progress on claim
			// trees. For instance we can run background jobs that look for expiring
			// claims and ensure the creation of claims in their next lifecycle phase.
			err = r.red.Sorted().Create().Score(posExp(inp[i].Lifecycle.Data), inp[i].ID.String(), float64(exp.Unix()))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if inp[i].Kind == "claim" {
			// Store every claim ID in a global sorted set, based on their time of
			// creation. This step ensures we can search for claims using pagination
			// in reverse chronolical order.
			err = r.red.Sorted().Create().Score(posLif(objectlabel.LifecycleCreated), inp[i].ID.String(), float64(inp[i].Created.Unix()))
			if err != nil {
				return tracer.Mask(err)
			}

			// We index all claim IDs per specified lifecycle phase, so that we can
			// search e.g. for all disputes.
			err = r.red.Sorted().Create().Score(posLif(inp[i].Lifecycle.Data), inp[i].ID.String(), float64(inp[i].Created.Unix()))
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
			err = r.red.Sorted().Create().Score(posOwnCom(inp[i].Owner, inp[i].Parent), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

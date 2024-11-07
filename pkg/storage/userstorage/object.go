package userstorage

import (
	"math"
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/round"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

const (
	Right     = 0
	Wrong     = 1
	Honest    = 2
	Dishonest = 3
	Abstained = 4
)

type Object struct {
	// Created is the time at which the user got created.
	Created time.Time `json:"created,omitempty"`
	// ID is the internal ID of the user being created.
	ID objectid.ID `json:"id,omitempty"`
	// Imag is the URL pointing to the user's profile picture.
	Image objectfield.String `json:"image,omitempty"`
	// Name is the user name.
	Name objectfield.String `json:"name,omitempty"`
	// Pointer is the user's mapping between note kinds and note pointers. The
	// pointers tracked here per notification topic keep track of the point up to
	// which users caught up with their notifications.
	Pointer map[string]string `json:"pointer,omitempty"`
	// Subject is the list of external subject claims mapped to the user being
	// created.
	Subject []string `json:"subject,omitempty"`
	// Summary is the collection of user specific metrics on competence and
	// integrity.
	//
	//     [0] the number of times this user has been right
	//     [1] the number of times this user has been wrong
	//     [2] the number of times this user has been honest
	//     [3] the number of times this user has been dishonest
	//     [4] the number of times this user has denied to vote
	//
	Summary []float64 `json:"summary,omitempty"`
}

func (o *Object) Competence() float64 {
	var rig float64
	{
		rig = o.Summary[Right]
	}

	if rig == 0 {
		return 0
	}

	var wro float64
	{
		wro = o.Summary[Wrong]
	}

	return rig / (rig + wro)
}

func (o *Object) Integrity() float64 {
	var hon float64
	{
		hon = o.Summary[Honest]
	}

	if hon == 0 {
		return 0
	}

	var dis float64
	var abs float64
	{
		dis = o.Summary[Dishonest]
		abs = o.Summary[Abstained] * 0.5
	}

	return math.Max(0, hon-dis-abs) + (hon / (hon + dis + abs))
}

func (o *Object) Reputation() float64 {
	return round.RoundP(o.Competence()*o.Integrity(), 5)
}

func (o *Object) Verify() error {
	{
		if o.Name.Data != "" && len(o.Name.Data) < 2 {
			return tracer.Maskf(UserNameLengthError, "%d", len(o.Name.Data))
		}
		if o.Name.Data != "" && len(o.Name.Data) > 30 {
			return tracer.Maskf(UserNameLengthError, "%d", len(o.Name.Data))
		}
	}

	{
		if len(o.Subject) != 1 || o.Subject[0] == "" {
			return tracer.Mask(UserSubjectEmptyError)
		}
	}

	{
		if len(o.Summary) != 5 {
			return tracer.Mask(UserSummaryError)
		}
	}

	return nil
}

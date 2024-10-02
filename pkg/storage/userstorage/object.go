package userstorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
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
	//
	Summary []int64 `json:"summary,omitempty"`
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

	return nil
}

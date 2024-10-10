package userstorage

import (
	"github.com/xh3b4sd/tracer"
)

// TODO add the given user IDs to the competence index
func (r *Redigo) UpdateCompetence(inp []*Object) error {
	// var err error
	return nil
}

// TODO add the given user IDs to the integrity index
func (r *Redigo) UpdateIntegrity(inp []*Object) error {
	// var err error
	return nil
}

func (r *Redigo) UpdateUser(inp []*Object) error {
	var err error

	for i := range inp {
		{
			err = inp[i].Verify()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Simple().Create().Element(useObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

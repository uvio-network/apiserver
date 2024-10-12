package userstorage

import (
	"errors"
	"time"

	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) CreateUser(inp *Object) (*Object, error) {
	var err error

	{
		err = inp.Verify()
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var out *Object
	{
		out, err = r.SearchSubject(inp.Subject[0])
		if errors.Is(err, SubjectClaimMappingError) {
			// The user does not appear to exist. So we initialize the user object.
			{
				inp.Created = now
				inp.ID = objectid.Random(objectid.Time(now))
				inp.Image.Time = now
				inp.Name.Time = now
			}

			// Here we create the mapping between external subject claim and internal
			// user ID.
			{
				err = r.red.Simple().Create().Element(useSub(inp.Subject[0]), inp.ID.String())
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// Now create the mapping between internal user ID and internal user
			// object.
			{
				err = r.red.Simple().Create().Element(useObj(inp.ID), musStr(inp))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			return inp, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		} else if out.Image.Data != inp.Image.Data || out.Name.Data != inp.Name.Data {
			{
				out.Image.Data = inp.Image.Data
				out.Image.Time = now
				out.Name.Data = inp.Name.Data
				out.Name.Time = now
			}

			{
				err = r.red.Simple().Create().Element(useObj(out.ID), musStr(out))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}
	}

	return out, nil
}

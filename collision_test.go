package sdf

import . "gopkg.in/check.v1"

type CollisionSuite struct{}

var _ = Suite(&CollisionSuite{})

func (s *CollisionSuite) TestCollisionValidation(c *C) {
	validCollision := []*Collision{
		&Collision{Name: "foo", Geometry: NewBox(Vec3{1, 1, 1})},
	}

	invalidCollision := []struct {
		Err       string
		Collision *Collision
	}{
		{
			Err:       "Missing name in sdf.Collision",
			Collision: &Collision{Geometry: NewBox(Vec3{1, 1, 1})},
		},
		{
			Err:       "Missing geometry in sdf.Collision",
			Collision: &Collision{Name: "foo"},
		},
	}

	for _, coll := range validCollision {
		c.Check(coll.Validate(), IsNil)
	}

	for _, d := range invalidCollision {
		c.Check(d.Collision.Validate(), ErrorMatches, d.Err)
	}

}

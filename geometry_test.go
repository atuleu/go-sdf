package sdf

import . "gopkg.in/check.v1"

type GeometrySuite struct{}

var _ = Suite(&GeometrySuite{})

func (s *GeometrySuite) TestValidation(c *C) {
	validGeometries := []*Geometry{
		&Geometry{},
		&Geometry{Box: &Box{Size: NewZeroVec3()}},
		&Geometry{Cylinder: &Cylinder{Radius: 1.0, Length: 1.0}},
		&Geometry{Heightmap: &Heightmap{}},
		&Geometry{Image: &Image{}},
		&Geometry{Mesh: &Mesh{}},
		&Geometry{Plane: &Plane{Normal: Vec3{1, 0, 0}, Size: Vec2{1, 1}}},
		&Geometry{Polyline: &Polyline{}},
		&Geometry{Sphere: &Sphere{Radius: 1}},
	}

	invalidGeometries := map[string]*Geometry{
		"Box,Cylinder": &Geometry{
			Box:      &Box{Size: NewZeroVec3()},
			Cylinder: &Cylinder{Radius: 1.0, Length: 1.0},
		},
	}

	for fields, g := range invalidGeometries {
		c.Check(g.Validate(), ErrorMatches, "Multiple geometry '"+fields+"' in sdf.Geometry")
	}

	for _, g := range validGeometries {
		c.Check(g.Validate(), IsNil)
	}

}

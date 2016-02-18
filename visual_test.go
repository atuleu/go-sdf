package sdf

import . "gopkg.in/check.v1"

type VisualSuite struct{}

var _ = Suite(&VisualSuite{})

func (s *VisualSuite) TestVisualValidation(c *C) {
	invalidData := []*Visual{
		&Visual{Name: ""},
	}

	validData := []*Visual{
		&Visual{Name: "foo"},
	}

	for _, v := range invalidData {
		c.Check(v.Validate(), ErrorMatches, "Missing name in sdf.Visual")
	}

	for _, v := range validData {
		c.Check(v.Validate(), IsNil)
	}

}

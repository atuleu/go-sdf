package sdf

import (
	. "gopkg.in/check.v1"
)

type LinkSuite struct{}

var _ = Suite(&LinkSuite{})

func (s *LinkSuite) TestLinkValidation(c *C) {
	invalidLink := []*Link{
		&Link{Name: ""},
	}

	validLink := []*Link{
		&Link{Name: "foo"},
	}

	for _, l := range invalidLink {
		c.Check(l.Validate(), ErrorMatches, "Missing name in sdf.Link")
	}

	for _, l := range validLink {
		c.Check(l.Validate(), IsNil)
	}

}

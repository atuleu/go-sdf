package sdf

import (
	. "gopkg.in/check.v1"
)

type FrameSuite struct{}

var _ = Suite(&FrameSuite{})

func (s *FrameSuite) TestFrameValidation(c *C) {
	invalidFrames := []*Frame{
		&Frame{},
		&Frame{Pose: &Pose{Values: []float64{0, 0, 0, 0, 0, 0}}},
	}
	validFrames := []*Frame{
		&Frame{Name: "foo"},
		&Frame{Name: "bar", Pose: &Pose{Values: []float64{0, 0, 0, 0, 0, 0}}},
	}
	for _, inc := range invalidFrames {
		c.Check(inc.Validate(), ErrorMatches, "Missing name in sdf.ModelFrame")
	}

	for _, inc := range validFrames {
		c.Check(inc.Validate(), IsNil)
	}
}

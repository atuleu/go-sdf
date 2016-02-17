package sdf

import (
	"encoding/xml"

	. "gopkg.in/check.v1"
)

type PoseSuite struct{}

var _ = Suite(&PoseSuite{})

func (s *PoseSuite) TestPoseHasFixedSize(c *C) {
	invalidPoses := []*Pose{
		&Pose{
			Frame:  "foo",
			Values: []float64{},
		},
		&Pose{
			Frame:  "",
			Values: []float64{1, 2, 3},
		},
	}
	validPoses := []*Pose{
		&Pose{
			Frame:  "bar",
			Values: []float64{1, 2, 3, 4, 5, 6},
		},
		&Pose{
			Values: []float64{0, 0, 0, 0, 0, 0},
		},
	}

	for _, p := range invalidPoses {
		c.Check(p.Validate(), ErrorMatches, `A sdf\.Pose should have 6 values \(Position\(x,y,z\), Rotation\(roll,pitch,yaw\)\), got .* elements`)
	}

	for _, p := range validPoses {
		c.Check(p.Validate(), IsNil)
	}
}

func (s *PoseSuite) TestXmlOutput(c *C) {
	data := map[string]*Pose{
		"<pose frame=\"foo\">1.000000 2.000000 3.000000 4.000000 5.000000 6.000000</pose>": &Pose{
			Frame:  "foo",
			Values: []float64{1, 2, 3, 4, 5, 6},
		},
	}

	for xmlValue, p := range data {
		d, err := xml.Marshal(p)
		c.Check(err, IsNil)
		c.Check(string(d), Equals, xmlValue)

		res := &Pose{}
		err = xml.Unmarshal([]byte(xmlValue), res)
		c.Check(err, IsNil)
		c.Check(res, DeepEquals, p)

	}
}

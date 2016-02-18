package sdf

import (
	"encoding/xml"

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

func (s *LinkSuite) TestLinkXml(c *C) {
	data := []struct {
		Xml string
		L   *Link
	}{
		{
			Xml: `<link name="foo"><pose>0.000000 0.000000 0.000000 0.000000 0.000000 0.000000</pose><inertial><inertia><ixx>1</ixx><ixy>0</ixy><ixz>0</ixz><iyy>1</iyy><iyz>0</iyz><izz>1</izz></inertia><mass>1</mass></inertial><self_collide>0</self_collide><kinematic>0</kinematic><gravity>1</gravity></link>`,
			L: &Link{
				Name: "foo",
				Pose: &Pose{Values: []float64{0, 0, 0, 0, 0, 0}},
				Inertial: &Inertial{
					Inertia: &Inertia{Xx: 1, Yy: 1, Zz: 1},
					Mass:    1,
				},
				Gravity:        true,
				NonSelfCollide: true,
				NonKinematic:   true,
			},
		},
	}

	for _, d := range data {
		xmlRes, err := xml.Marshal(d.L)
		c.Check(err, IsNil)
		c.Check(string(xmlRes), Equals, d.Xml)

		lRes := &Link{}
		err = xml.Unmarshal([]byte(d.Xml), lRes)
		c.Check(err, IsNil)
		c.Check(lRes, DeepEquals, d.L)
		c.Check(lRes.Validate(), IsNil)

	}

}

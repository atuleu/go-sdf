package sdf

import (
	"encoding/xml"

	. "gopkg.in/check.v1"
)

type PoseSuite struct{}

var _ = Suite(&PoseSuite{})

func (s *PoseSuite) TestXmlOutput(c *C) {
	data := map[string]*Pose{
		"<pose frame=\"foo\">1 2 3 4 5 6</pose>": &Pose{
			Vec6:  Vec6{1, 2, 3, 4, 5, 6},
			Frame: "foo",
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

package sdf

import (
	"encoding/xml"

	. "gopkg.in/check.v1"
)

type VisualSuite struct{}

var _ = Suite(&VisualSuite{})

func (s *VisualSuite) TestVisualValidation(c *C) {
	invalidData := []struct {
		Err    string
		Visual *Visual
	}{
		{
			Err:    "Missing name in sdf.Visual",
			Visual: &Visual{Name: "", Geometry: NewBox(Vec3{1, 1, 1})},
		},
		{
			Err:    "Missing sdf.Geometry in sdf.Visual",
			Visual: &Visual{Name: "foo"},
		},
	}

	validData := []*Visual{
		&Visual{Name: "foo", Geometry: NewBox(Vec3{1, 1, 1})},
	}

	for _, d := range invalidData {
		c.Check(d.Visual.Validate(), ErrorMatches, d.Err)
	}

	for _, v := range validData {
		c.Check(v.Validate(), IsNil)
	}

}

func (s *VisualSuite) TestVisualXMLOutput(c *C) {
	data := []struct {
		Xml    string
		Visual *Visual
	}{
		{
			Visual: NewVisual("foo", NewBox(Vec3{1, 1, 1})),
			Xml:    "<visual name=\"foo\"><geometry><box><size>1 1 1</size></box></geometry></visual>",
		},
	}

	for _, d := range data {
		resXml, err := xml.Marshal(d.Visual)
		c.Check(err, IsNil)
		c.Check(string(resXml), DeepEquals, d.Xml)

		resVisual := &Visual{}
		err = xml.Unmarshal([]byte(d.Xml), resVisual)
		c.Check(err, IsNil)
		c.Check(resVisual, DeepEquals, d.Visual)
		c.Check(resVisual.Validate(), IsNil)
	}
}

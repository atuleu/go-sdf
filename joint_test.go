package sdf

import (
	"encoding/xml"

	. "gopkg.in/check.v1"
)

type JointSuite struct{}

var _ = Suite(&JointSuite{})

func (s *JointSuite) TestJointValidation(c *C) {
	type InvalidData struct {
		Err   string
		Joint *Joint
	}
	invalidJoint := []InvalidData{
		{
			Err:   "Missing name in sdf.Joint",
			Joint: &Joint{Type: JOINT_BALL, Parent: "a", Child: "b"},
		},
		{
			Err:   "Invalid sdf.Joint.Type ''",
			Joint: &Joint{Name: "foo", Type: "", Parent: "a", Child: "b"},
		},
		{
			Err:   "Invalid sdf.Joint.Type 'does_not_exist_as_a_type'",
			Joint: &Joint{Name: "foo", Type: "does_not_exist_as_a_type", Parent: "a", Child: "b"},
		},
	}

	validJoint := []*Joint{}
	for t, _ := range availableJointType {
		invalidJoint = append(invalidJoint,
			InvalidData{
				Err:   "Missing name in sdf.Joint",
				Joint: &Joint{Name: "", Type: t},
			})
		invalidJoint = append(invalidJoint,
			InvalidData{
				Err:   "Missing parent in sdf.Joint",
				Joint: &Joint{Name: "foo", Type: t},
			})
		invalidJoint = append(invalidJoint,
			InvalidData{
				Err:   "Missing child in sdf.Joint",
				Joint: &Joint{Name: "foo", Type: t, Parent: "a"},
			})

		validJoint = append(validJoint, &Joint{Name: "foo", Type: t, Parent: "a", Child: "b"})
	}

	for _, j := range validJoint {
		c.Check(j.Validate(), IsNil)
	}

	for _, d := range invalidJoint {
		c.Check(d.Joint.Validate(), ErrorMatches, d.Err)
	}

}

func (s *JointSuite) TestJoinXMLOutput(c *C) {
	type TestData struct {
		Xml   string
		Joint *Joint
	}

	data := []TestData{
		{
			Xml: `<joint name="bar_12_joint" type="revolute"><parent>link_1</parent><child>link_2</child><pose>0 0.5 0 0 0 0</pose><axis><xyz>0 0 1</xyz></axis></joint>`,
			Joint: &Joint{
				Name:   "bar_12_joint",
				Type:   JOINT_REVOLUTE,
				Parent: "link_1",
				Child:  "link_2",
				Pose:   NewPose(0, 0.5, 0, 0, 0, 0),
				Axis:   NewAxis(Vec3{0, 0, 1}),
			},
		},
	}

	for _, d := range data {
		xmlRes, err := xml.Marshal(d.Joint)
		c.Check(err, IsNil)
		c.Check(string(xmlRes), Equals, d.Xml)

		j := &Joint{}
		err = xml.Unmarshal([]byte(d.Xml), j)
		c.Check(err, IsNil)
		c.Check(j.Validate(), IsNil)
		c.Check(j, DeepEquals, d.Joint)

	}

}

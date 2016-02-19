package sdf

import . "gopkg.in/check.v1"

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
			Joint: &Joint{},
		},
		{
			Err:   "Invalid sdf.Joint.Type ''",
			Joint: &Joint{Name: "foo"},
		},
		{
			Err:   "Invalid sdf.Joint.Type 'does_not_exist_as_a_type'",
			Joint: &Joint{Name: "foo", Type: "does_not_exist_as_a_type"},
		},
	}

	validJoint := []*Joint{}
	for t, _ := range availableJointType {
		invalidJoint = append(invalidJoint,
			InvalidData{
				Err:   "Missing name in sdf.Joint",
				Joint: &Joint{Name: "", Type: t},
			})
		validJoint = append(validJoint, &Joint{Name: "foo", Type: t})
	}

	for _, j := range validJoint {
		c.Check(j.Validate(), IsNil)
	}

	for _, d := range invalidJoint {
		c.Check(d.Joint.Validate(), ErrorMatches, d.Err)
	}

}

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

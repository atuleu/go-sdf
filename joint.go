package sdf

import "fmt"

const (
	JOINT_REVOLUTE  = "revolute"
	JOINT_REVOLUTE2 = "revolute2"
	JOINT_GEARBOX   = "gearbox"
	JOINT_SCREW     = "screw"
	JOINT_BALL      = "ball"
	JOINT_PRISMATIC = "prismatic"
	JOINT_FIXED     = "fixed"
	JOINT_UNIVERSAL = "universal"
)

type Joint struct {
	XMLName struct{} `xml:"joint"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type"`
}

var availableJointType = map[string]struct{}{
	JOINT_REVOLUTE:  struct{}{},
	JOINT_REVOLUTE2: struct{}{},
	JOINT_GEARBOX:   struct{}{},
	JOINT_SCREW:     struct{}{},
	JOINT_BALL:      struct{}{},
	JOINT_PRISMATIC: struct{}{},
	JOINT_FIXED:     struct{}{},
	JOINT_UNIVERSAL: struct{}{},
}

func (j *Joint) Validate() error {
	if len(j.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.Joint")
	}

	if _, ok := availableJointType[j.Type]; ok == false {
		return fmt.Errorf("Invalid sdf.Joint.Type '%s'", j.Type)
	}

	return nil
}

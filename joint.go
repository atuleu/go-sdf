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

type AxisDynamic struct {
	XMLName         struct{} `xml:"dynamics"`
	Damping         float64  `xml:"damping,omitempty"`
	Friction        float64  `xml:"friction,omitempty"`
	SpringReference float64  `xml:"spring_reference"`
	SpringStiffness float64  `xml:"spring_stiffness"`
}

type AxisLimit struct {
	XMLName     struct{} `xml:"limit"`
	Lower       float64  `xml:"lower,omitempty"`
	Upper       float64  `xml:"upper,omitempty"`
	Effort      float64  `xml:"effort,omitempty"`
	Velocity    float64  `xml:"velocity,omitempty"`
	Stiffness   float64  `xml:"stiffness,omitempty"`
	Dissipation float64  `xml:"dissipation,omitempty"`
}

type Axis struct {
	Xyz            Vec3         `xml:"xyz"`
	UseParentFrame Bool         `xml:"use_parent_model_frame"`
	Dynamics       *AxisDynamic `xml:"dynamics,omitempty"`
	Limit          *AxisLimit   `xml:"limit,omitempty"`
}

func NewAxis(value Vec3) *Axis {
	return &Axis{Xyz: value}
}

type JointPhysicsODE struct {
	XMLName             struct{} `xml:"ode"`
	CFMDamping          Bool     `xml:"cfm_damping"`
	ImplicitSpingDamper Bool     `xml:"implicit_spring_damper"`
	FudgeFactor         float64  `xml:"fudge_factor,omitempty"`
	Cfm                 float64  `xml:"cfm,omitempty"`
	Erp                 float64  `xml:"erp,omitempty"`
}

type JointPhysicsSimbody struct {
	XMLName         struct{} `xml:"simbody"`
	MustBeLoopJoint Bool     `xml:"must_be_loop_joint"`
}

type JointPhysics struct {
	XMLName struct{}             `xml:"physics"`
	Simbody *JointPhysicsSimbody `xml:"simbody"`
	ODE     *JointPhysicsODE     `xml:"ode"`
}

type Joint struct {
	XMLName struct{} `xml:"joint"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Parent  string   `xml:"parent"`
	Child   string   `xml:"child"`

	Pose   *Pose    `xml:"pose"`
	Frames []*Frame `xml:"frame"`

	Axis  *Axis `xml:"axis,omitempty"`
	Axis2 *Axis `xml:"axis2,omitempty"`

	GearboxRatio         float64 `xml:"gearbox_ratio,omitempty"`
	GearboxReferenceBody string  `xml:"gearbox_reference_body,omitempty"`

	ThreadPitch float64 `xml:"thread_pitch,omitempty"`

	Physics *JointPhysics `xml:"physics"`
}

func (j *Joint) Validate() error {
	if len(j.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.Joint")
	}

	if _, ok := availableJointType[j.Type]; ok == false {
		return fmt.Errorf("Invalid sdf.Joint.Type '%s'", j.Type)
	}

	if len(j.Parent) == 0 {
		return fmt.Errorf("Missing parent in sdf.Joint")
	}

	if len(j.Child) == 0 {
		return fmt.Errorf("Missing child in sdf.Joint")
	}

	return nil
}

func NewRevoluteJoint(name string, axis Vec3) *Joint {
	return &Joint{
		Name: name,
		Type: JOINT_REVOLUTE,
		Axis: NewAxis(axis),
	}
}

func NewPrismaticJoint(name string, axis Vec3) *Joint {
	return &Joint{
		Name: name,
		Type: JOINT_PRISMATIC,
		Axis: NewAxis(axis),
	}
}

func NewGearboxJoint(name string, axis Vec3, ratio float64, ref string) *Joint {
	return &Joint{
		Name:                 name,
		Type:                 JOINT_GEARBOX,
		Axis:                 NewAxis(axis),
		GearboxRatio:         ratio,
		GearboxReferenceBody: ref,
	}
}

func NewScrewJoint(name string, axis Vec3, threadPitch float64) *Joint {
	return &Joint{
		Name:        name,
		Type:        JOINT_SCREW,
		Axis:        NewAxis(axis),
		ThreadPitch: threadPitch,
	}
}

func NewRevolute2Joint(name string, axis, axis2 Vec3) *Joint {
	return &Joint{
		Name:  name,
		Type:  JOINT_REVOLUTE2,
		Axis:  NewAxis(axis),
		Axis2: NewAxis(axis2),
	}
}

func NewBallJoint(name string) *Joint {
	return &Joint{
		Name: name,
		Type: JOINT_BALL,
	}
}

func NewFixedJoint(name string) *Joint {
	return &Joint{
		Name: name,
		Type: JOINT_FIXED,
	}
}

func NewUniversalJoint(name string, axis Vec3) *Joint {
	return &Joint{
		Name: name,
		Type: JOINT_UNIVERSAL,
		Axis: NewAxis(axis),
	}
}

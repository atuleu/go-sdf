package sdf

import (
	"encoding/xml"
	"fmt"
)

type VelocityDecay struct {
	XMLName xml.Name `xml:"velocity_decay"`
	Linear  float64  `xml:"linear"`
	Angular float64  `xml:"angular"`
}

type Inertia struct {
	Xx float64 `xml:"ixx"`
	Xy float64 `xml:"ixy"`
	Xz float64 `xml:"ixz"`
	Yy float64 `xml:"iyy"`
	Yz float64 `xml:"iyz"`
	Zz float64 `xml:"izz"`
}

type Inertial struct {
	XMLName struct{} `xml:"inertial"`
	Inertia *Inertia `xml:"inertia"`
	Mass    float64  `xml:"mass"`
	Frame   *Frame   `xml:",omitempty"`
	Pose    *Pose    `xml:",omitempty"`
}

type Link struct {
	XMLName        struct{}       `xml:"link"`
	Name           string         `xml:"name,attr"`
	Pose           *Pose          `xml:"pose,omitempty"`
	Inertial       *Inertial      `xml:",omitempty"`
	NonSelfCollide InvertedBool   `xml:"self_collide"`
	NonKinematic   InvertedBool   `xml:"kinematic"`
	Gravity        Bool           `xml:"gravity"`
	MustBeBaseLink Bool           `xml:"must_be_base_link,omitempty"`
	VelocityDecay  *VelocityDecay `xml:",omitempty"`
	Frames         []*Frame       `xml:"frame"`
	Visuals        []*Visual
	Collisions     []*Collision
}

func NewLink(name string) *Link {
	return &Link{
		Name:           name,
		NonSelfCollide: true,
		NonKinematic:   true,
	}
}

func (l *Link) Validate() error {
	if len(l.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.Link")
	}
	return nil
}

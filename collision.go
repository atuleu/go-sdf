package sdf

import "fmt"

type Surface struct {
	XMLName struct{} `xml:"surface"`
}

type Collision struct {
	XMLName     struct{} `xml:"collision"`
	Name        string   `xml:"name,attr"`
	MaxContacts int      `xml:"max_contacts,omitempty"`
	Frames      []*Frame
	Pose        *Pose `xml:",omitempty"`
	Geometry    *Geometry
	Surface     *Surface `xml:",omitempty"`
}

func (c *Collision) Validate() error {
	if len(c.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.Collision")
	}
	if c.Geometry == nil {
		return fmt.Errorf("Missing geometry in sdf.Collision")
	}

	return nil
}

func NewCollision(name string, g *Geometry) *Collision {
	return &Collision{Name: name, Geometry: g}
}

package sdf

import "fmt"

type Visual struct {
	XMLName      struct{} `xml:"visual"`
	Name         string   `xml:"name,attr"`
	CastShadows  Bool     `xml:"cast_shadows"` // TODO check omitempty
	Transparency float64  `xml:"transparency,omitempty"`
	//TODO add Meta
	Frames   []*Frame
	Pose     *Pose
	Material *Material `xml:"material,omitempty"`
	Geometry *Geometry `xml:"geometry,omitempty"`
}

func (v *Visual) Validate() error {
	if len(v.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.Visual")
	}
	if v.Geometry == nil {
		return fmt.Errorf("Missing sdf.Geometry in sdf.Visual")
	}
	return nil
}

func NewVisual(name string, g *Geometry) *Visual {
	return &Visual{Name: name, Geometry: g}
}

package sdf

import "fmt"

type Material struct {
	XMLName struct{} `xml:"material"`
}

type Visual struct {
	XMLName      struct{} `xml:"visual"`
	Name         string   `xml:"name"`
	CastShadows  Bool     `xml:"cast_shadows"` // TODO check omitempty
	Transparency float64  `xml:"transparency"`
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
	return nil
}

func NewVisual(name string) *Visual {
	return &Visual{Name: name}
}

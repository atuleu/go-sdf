package sdf

import (
	"encoding/xml"
	"fmt"
)

type Pose struct {
	Frame  string
	Values []float64
}

func (p *Pose) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := p.Validate(); err != nil {
		return err
	}
	start.Name.Local = "pose"
	if len(p.Frame) > 0 {
		start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Space: "", Local: "frame"}, Value: p.Frame}}
	}
	str := fmt.Sprintf("%f %f %f %f %f %f", p.Values[0], p.Values[1], p.Values[2], p.Values[3], p.Values[4], p.Values[5])
	if err := e.EncodeElement(str, start); err != nil {
		return err
	}

	return nil
}

func (p *Pose) Validate() error {
	if len(p.Values) != 6 {
		return fmt.Errorf("A sdf.Pose should have 6 values (Position(x,y,z), Rotation(roll,pitch,yaw)), got %d elements", len(p.Values))
	}

	return nil
}

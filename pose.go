package sdf

import (
	"encoding/xml"
	"fmt"
)

type Pose struct {
	Frame  string
	Values []float64
}

type golangPose struct {
	Frame  string `xml:"frame,attr,omitempty"`
	Values string `xml:",chardata"`
}

func (p *Pose) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := p.Validate(); err != nil {
		return err
	}
	start.Name.Local = "pose"
	pp := &golangPose{
		Frame:  p.Frame,
		Values: fmt.Sprintf("%f %f %f %f %f %f", p.Values[0], p.Values[1], p.Values[2], p.Values[3], p.Values[4], p.Values[5]),
	}

	return e.EncodeElement(pp, start)
}

func (p *Pose) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	pp := &golangPose{}
	if err := d.DecodeElement(pp, &start); err != nil {
		return err
	}

	p.Frame = pp.Frame
	p.Values = make([]float64, 6)

	var x, y, z, roll, pitch, yaw float64

	_, err := fmt.Sscanf(pp.Values, "%f %f %f %f %f %f", &x, &y, &z, &roll, &pitch, &yaw)

	p.Values[0] = x
	p.Values[1] = y
	p.Values[2] = z
	p.Values[3] = roll
	p.Values[4] = pitch
	p.Values[5] = yaw

	return err
}

func (p *Pose) Validate() error {
	if len(p.Values) != 6 {
		return fmt.Errorf("A sdf.Pose should have 6 values (Position(x,y,z), Rotation(roll,pitch,yaw)), got %d elements", len(p.Values))
	}

	return nil
}

package sdf

import "encoding/xml"

type Pose struct {
	XMLName struct{} `xml:"pose"`
	Vec6
	Frame string `xml:"frame,attr"`
}

func NewZeroPose() *Pose {
	return &Pose{
		Vec6:  NewZeroVec6(),
		Frame: "",
	}
}

func NewPose(x, y, z, roll, pitch, yaw float64) *Pose {
	return &Pose{
		Vec6:  Vec6{x, y, z, roll, pitch, yaw},
		Frame: "",
	}
}

func (p *Pose) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pose"
	if len(p.Frame) != 0 {
		start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "frame"}, Value: p.Frame}}
	}
	return e.EncodeElement(p.Vec6, start)
}

func (p *Pose) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data Vec6
	if err := d.DecodeElement(&data, &start); err != nil {
		return err
	}
	(*p).Vec6 = data
	p.Frame = ""
	for _, a := range start.Attr {
		if a.Name.Local != "frame" {
			continue
		}
		p.Frame = a.Value
	}
	return nil
}

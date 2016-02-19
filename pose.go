package sdf

import (
	"encoding/xml"
	"math"

	"github.com/skelterjohn/go.matrix"
)

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

func yawMatrix(yaw float64) *matrix.DenseMatrix {
	return matrix.MakeDenseMatrix([]float64{math.Cos(yaw), -math.Sin(yaw), 0, math.Sin(yaw), math.Cos(yaw), 0, 0, 0, 1}, 3, 3)
}

func pitchMatrix(pitch float64) *matrix.DenseMatrix {
	return matrix.MakeDenseMatrix([]float64{math.Cos(pitch), 0, math.Sin(pitch), 0, 1, 0, -math.Sin(pitch), 0, math.Cos(pitch)}, 3, 3)
}

func rollMatrix(roll float64) *matrix.DenseMatrix {
	return matrix.MakeDenseMatrix([]float64{1, 0, 0, 0, math.Cos(roll), -math.Sin(roll), 0, math.Sin(roll), math.Cos(roll)}, 3, 3)
}

func (p *Pose) NewRelativePose(x, y, z, roll, pitch, yaw float64) *Pose {
	vec := matrix.Product(
		matrix.Product(
			matrix.Product(yawMatrix(p.Vec6[5]), pitchMatrix(p.Vec6[4])),
			rollMatrix(p.Vec6[3])),
		matrix.MakeDenseMatrix([]float64{x, y, z}, 3, 1))

	return NewPose(p.Vec6[0]+vec.Array()[0], p.Vec6[1]+vec.Array()[1], p.Vec6[2]+vec.Array()[2],
		p.Vec6[3]+roll, p.Vec6[4]+pitch, p.Vec6[5]+yaw)
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

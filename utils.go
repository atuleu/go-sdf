package sdf

import (
	"encoding/xml"
	"fmt"
)

type Bool bool

type InvertedBool bool

func (b Bool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if bool(b) == false {
		return nil
	}

	asInt := 1
	return e.EncodeElement(asInt, start)
}

func (b *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	asInt := 0
	if err := d.DecodeElement(&asInt, &start); err != nil {
		return err
	}
	if asInt != 0 {
		*b = Bool(true)
	}
	return nil
}

func (b InvertedBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if bool(b) == false {
		return nil
	}

	asInt := 0
	return e.EncodeElement(asInt, start)
}

func (b *InvertedBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*b = InvertedBool(false)
	asInt := 0
	if err := d.DecodeElement(&asInt, &start); err != nil {
		return err
	}
	if asInt == 0 {
		*b = InvertedBool(true)
	}
	return nil
}

type Vec2 [2]float64

func NewZeroVec2() Vec2 {
	return Vec2{0, 0}
}

type Vec3 [3]float64

func NewZeroVec3() Vec3 {
	return Vec3{0, 0, 0}
}

var UnitX = Vec3{1, 0, 0}
var UnitY = Vec3{0, 1, 0}
var UnitZ = Vec3{0, 0, 1}

type Vec4 [4]float64

func NewZeroVec4() Vec4 {
	return Vec4{0, 0, 0, 0}
}

type Vec6 [6]float64

func NewZeroVec6() Vec6 {
	return Vec6{0, 0, 0, 0, 0, 0}
}

func (v Vec2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	str := fmt.Sprintf("%v", v)
	return e.EncodeElement(str[1:len(str)-1], start)
}

func (v Vec3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	str := fmt.Sprintf("%v", v)
	return e.EncodeElement(str[1:len(str)-1], start)
}

func (v Vec4) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	str := fmt.Sprintf("%v", v)
	return e.EncodeElement(str[1:len(str)-1], start)
}

func (v Vec6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	str := fmt.Sprintf("%v", v)
	return e.EncodeElement(str[1:len(str)-1], start)
}

func (v *Vec2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	str := ""
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	_, err := fmt.Sscanf(str, "%f %f", &v[0], &v[1])
	return err
}

func (v *Vec3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	str := ""
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	_, err := fmt.Sscanf(str, "%f %f %f", &v[0], &v[1], &v[2])
	return err
}

func (v *Vec4) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	str := ""
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	_, err := fmt.Sscanf(str, "%f %f %f %f", &v[0], &v[1], &v[2], &v[3])
	return err
}

func (v *Vec6) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	str := ""
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	_, err := fmt.Sscanf(str, "%f %f %f %f %f %f", &v[0], &v[1], &v[2], &v[3], &v[4], &v[5])
	return err
}

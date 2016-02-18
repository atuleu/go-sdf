package sdf

import "encoding/xml"

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

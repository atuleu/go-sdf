package sdf

import (
	"encoding/xml"
	"fmt"
)

type Link struct {
	XMLName xml.Name `xml:"link"`
	Name    string   `xml:"name,attr"`
}

func (l *Link) Validate() error {
	if len(l.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.Link")
	}
	return nil
}

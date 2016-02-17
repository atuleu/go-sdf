package sdf

import (
	"encoding/xml"
	"fmt"
	"regexp"
)

type World struct{}
type Model struct{}
type Actor struct{}
type Light struct{}

type Sdf struct {
	XMLName xml.Name `xml:"sdf"`
	Version string   `xml:"version,attr"`
	World   *World   `xml:"world,omitempty"`
	Model   *Model   `xml:"model,omitempty"`
	Actor   *Actor   `xml:"actor,omitempty"`
	Light   *Light   `xml:"light,omitempty"`
}

var versionRegexp = regexp.MustCompile(`1\.(4|5|6)`)

func (s *Sdf) Validate() error {
	if versionRegexp.MatchString(s.Version) == false {
		return fmt.Errorf("Invalid sdf version '%s'", s.Version)
	}

	return nil
}

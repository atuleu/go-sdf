package sdf

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type ModelInclude struct {
	XMLName xml.Name `xml:"include"`
	URI     string   `xml:"uri"`
	Pose    *Pose    `xml:"pose,omitempty"`
	Name    string   `xml:"name,omitempty"`
	Static  bool     `xml:"static,omitempty"`
}

func (inc *ModelInclude) Validate() error {
	if len(inc.URI) == 0 {
		return fmt.Errorf("Missing URI in sdf.ModelInclude")
	}
	return nil
}

type ModelFrame struct {
	Name string `xml:"frame"`
	Pose *Pose  `xml:"pose,omitempty"`
}

func (f *ModelFrame) Validate() error {
	if len(f.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.ModelFrame")
	}
	return nil
}

type Link struct{}

type Joint struct{}

type ModelPlugin struct {
	XMLName  xml.Name `xml:"plugin"`
	Name     string   `xml:"name,attr"`
	Filename string   `xml:"filename,attr`
}

func (p *ModelPlugin) Validate() error {
	missing := make([]string, 0, 2)
	if len(p.Name) == 0 {
		missing = append(missing, "name")
	}
	if len(p.Filename) == 0 {
		missing = append(missing, "filename")
	}
	if len(missing) == 0 {
		return nil
	}
	return fmt.Errorf("Missing %s in sdf.ModelPlugin", strings.Join(missing, ","))
}

type ModelGripper struct {
	XMLName xml.Name `xml:"gripper"`
	Name    string   `xml:"name,attr"`
}

type Model struct {
	XMLName          xml.Name `xml:"model"`
	Name             string   `xml:"name,attr"`
	Static           bool     `xml:"static,omitempty"`
	SelfCollide      bool     `xml:"self_collide,omitempty"`
	AllowAutoDisable bool     `xml:"allow_auto_disable,omitempty"`
	Includes         []*ModelInclude
	Models           []*Model
	EnableWind       bool        `xml:"enable_wind,omitempty"`
	Frame            *ModelFrame `xml:"frame,omitempty"`
	Pose             *Pose       `xml:"pose,omitempty"`
	Links            []*Link
	Plugins          []*ModelPlugin
	Grippers         []*ModelGripper
}

package sdf

import "encoding/xml"

type Pose struct {
	Frame string `xml:"frame,attr,omitempty"`
}

type ModelInclude struct {
	XMLName xml.Name `xml:"include"`
	URI     string   `xml:"uri"`
	Pose    *Pose    `xml:"pose,omitempty"`
	Name    string   `xml:"name,omitempty"`
	Static  bool     `xml:"static,omitempty"`
}

type ModelFrame struct {
	Name string `xml:"frame"`
	Pose *Pose  `xml:"pose,omitempty"`
}

type Link struct{}

type Joint struct{}

type ModelPlugin struct {
	XMLName  xml.Name `xml:"plugin"`
	Name     string   `xml:"name,attr"`
	Filename string   `xml:"filename,attr`
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

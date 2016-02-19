package sdf

import "fmt"

type Script struct {
	XMLName struct{} `xml:"script"`
	Name    string   `xml:"name"`
	URI     string   `xml:"uri,omitempty"`
}

func (s *Script) Validate() error {
	if len(s.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.Script")
	}
	return nil
}

func (s *Script) NewScript(name string) *Script {
	return &Script{Name: name}
}

const (
	SHD_VERTEX                  = "vertex"
	SHD_PIXEL                   = "pixel"
	SHD_NORMAL_MAP_OBJECTSPACE  = "normal_map_objectspace"
	SHD_NORMAL_MAP_TANGENTSPACE = "normal_map_tangentspace"
)

type Shader struct {
	XMLName   struct{} `xml:"shader"`
	Type      string   `xml:"type,attr"`
	NormalMap string   `xml:"normal_map"`
}

type Material struct {
	XMLName  struct{} `xml:"material"`
	Script   *Script  `xml:"script"`
	Shader   *Shader  `xml:"shader"`
	Lighting Bool     `xml:"lighting"`
	Ambient  Vec4     `xml:"ambient,omitempty"`
	Diffuse  Vec4     `xml:"diffuse,omitempty"`
	Specular Vec4     `xml:"specular,omitempty"`
	Emissive Vec4     `xml:"emmissive,omitempty"`
}

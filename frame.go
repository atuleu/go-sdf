package sdf

import "fmt"

type Frame struct {
	Name string `xml:"frame"`
	Pose *Pose  `xml:"pose,omitempty"`
}

func (f *Frame) Validate() error {
	if len(f.Name) == 0 {
		return fmt.Errorf("Missing name in sdf.ModelFrame")
	}
	return nil
}

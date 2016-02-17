package sdf

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type SdfSuite struct{}

var _ = Suite(&SdfSuite{})

func (s *SdfSuite) TestVersionIsRequired(c *C) {
	invalidVersion := []string{"", "1.2", "3.0.4~rc1"}
	for _, v := range invalidVersion {
		sdf := &Sdf{Version: v}
		c.Check(sdf.Validate(), ErrorMatches, "Invalid sdf version '"+v+"'")
	}
	validVersion := []string{"1.4", "1.5", "1.6"}
	for _, v := range validVersion {
		sdf := &Sdf{Version: v}
		c.Check(sdf.Validate(), IsNil)
	}
}

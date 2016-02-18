package sdf

import (
	. "gopkg.in/check.v1"
)

type ModelSuite struct{}

var _ = Suite(&ModelSuite{})

func (s *ModelSuite) TestIncludeValidation(c *C) {
	invalidIncludes := []*ModelInclude{
		&ModelInclude{},
		&ModelInclude{Static: true},
		&ModelInclude{Name: "foo"},
		&ModelInclude{Pose: NewZeroPose()},
	}

	validIncludes := []*ModelInclude{
		&ModelInclude{URI: "./foo"},
	}
	for _, inc := range invalidIncludes {
		c.Check(inc.Validate(), ErrorMatches, "Missing URI in sdf.ModelInclude")
	}

	for _, inc := range validIncludes {
		c.Check(inc.Validate(), IsNil)
	}

}

func (s *ModelSuite) TestPluginValidation(c *C) {
	invalidPlugins := map[string]*ModelPlugin{
		"name,filename": &ModelPlugin{},
		"filename":      &ModelPlugin{Name: "foo"},
		"name":          &ModelPlugin{Filename: "libfoo.so"},
	}

	validPlugins := []*ModelPlugin{
		&ModelPlugin{Name: "foo", Filename: "libfoo.so"},
	}

	for missing, inc := range invalidPlugins {
		c.Check(inc.Validate(), ErrorMatches, "Missing "+missing+" in sdf.ModelPlugin")
	}

	for _, inc := range validPlugins {
		c.Check(inc.Validate(), IsNil)
	}

}

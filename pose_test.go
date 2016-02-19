package sdf

import (
	"encoding/xml"
	"math"

	"github.com/skelterjohn/go.matrix"
	. "gopkg.in/check.v1"
)

type PoseSuite struct{}

var _ = Suite(&PoseSuite{})

func (s *PoseSuite) TestXmlOutput(c *C) {
	data := map[string]*Pose{
		"<pose frame=\"foo\">1 2 3 4 5 6</pose>": &Pose{
			Vec6:  Vec6{1, 2, 3, 4, 5, 6},
			Frame: "foo",
		},
	}

	for xmlValue, p := range data {
		d, err := xml.Marshal(p)
		c.Check(err, IsNil)
		c.Check(string(d), Equals, xmlValue)

		res := &Pose{}
		err = xml.Unmarshal([]byte(xmlValue), res)
		c.Check(err, IsNil)
		c.Check(res, DeepEquals, p)

	}
}

func rad2deg(rad float64) float64 {
	return rad * 180 / math.Pi
}

func deg2rad(rad float64) float64 {
	return rad * math.Pi / 180
}

func (s *PoseSuite) TestRelativePoseCreation(c *C) {
	type TestData struct {
		Base, Relative, Expected *Pose
	}

	data := []TestData{
		{
			Base:     NewZeroPose(),
			Relative: NewZeroPose(),
			Expected: NewZeroPose(),
		},
		{
			Base:     NewZeroPose(),
			Relative: NewPose(1, 2, 3, 4, 5, 6),
			Expected: NewPose(1, 2, 3, 4, 5, 6),
		},
		{
			Base:     NewPose(1, 2, 3, 4, 5, 6),
			Relative: NewZeroPose(),
			Expected: NewPose(1, 2, 3, 4, 5, 6),
		},
		{
			Base:     NewPose(1, 0, 0, deg2rad(45), 0, 0),
			Relative: NewPose(1, 0, 0, deg2rad(45), 0, 0),
			Expected: NewPose(2, 0, 0, deg2rad(90), 0, 0),
		},
		{
			Base:     NewPose(1, 0, 0, deg2rad(45), 0, 0),
			Relative: NewPose(1, 1, 1, deg2rad(45), 0, 0),
			Expected: NewPose(2, 0, math.Sqrt(2), deg2rad(90), 0, 0),
		},
		{
			Base:     NewPose(1, 0, 0, 0, deg2rad(45), 0),
			Relative: NewPose(1, 1, 1, deg2rad(45), 0, 0),
			Expected: NewPose(1+math.Sqrt(2), 1, 0, deg2rad(45), deg2rad(45), 0),
		},
		{
			Base:     NewPose(1, 0, 0, 0, 0, deg2rad(45)),
			Relative: NewPose(1, 1, 1, deg2rad(45), 0, 0),
			Expected: NewPose(1, math.Sqrt(2), 1, deg2rad(45), 0, deg2rad(45)),
		},
		{
			Base:     NewPose(1, 0, 0, deg2rad(90), deg2rad(90), 0),
			Relative: NewPose(1, 0, 0, 0, 0, 0),
			Expected: NewPose(1, 0, -1, deg2rad(90), deg2rad(90), 0),
		},
	}

	for _, d := range data {
		result := d.Base.NewRelativePose(d.Relative.Vec6[0], d.Relative.Vec6[1], d.Relative.Vec6[2],
			d.Relative.Vec6[3], d.Relative.Vec6[4], d.Relative.Vec6[5])
		c.Check(matrix.ApproxEquals(matrix.MakeDenseMatrix(result.Vec6[:], 6, 1),
			matrix.MakeDenseMatrix(d.Expected.Vec6[:], 6, 1),
			1e-10),
			Equals, true,
			Commentf("When computing %v relative to %v, expected, %v got %v", d.Relative.Vec6, d.Base.Vec6, d.Expected.Vec6, result.Vec6))
	}

}

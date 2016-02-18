package sdf

import (
	"fmt"
	"reflect"
	"strings"
)

type Box struct {
	Size Vec3 `xml:"size"`
}

type Cylinder struct {
	Radius float64 `xml:"radius"`
	Length float64 `xml:"length"`
}

type Heightmap struct {
	URI  string `xml:"uri"`
	Size *Vec3  `xml:"size,omitempty"`
	Pos  *Vec3  `xml:"pos,omitempty"`
}

type Image struct{}

type Mesh struct{}

type Plane struct {
	Normal Vec3 `xml:"normal"`
	Size   Vec2 `xml:"size"`
}

type Polyline struct{}

type Sphere struct {
	Radius float64 `xml:"radius"`
}

type Geometry struct {
	XMLName   struct{}   `xml:"geometry"`
	Box       *Box       `xml:"box,omitempty"`
	Cylinder  *Cylinder  `xml:"cylinder,omitempty"`
	Heightmap *Heightmap `xml:"heightmap,omitempty"`
	Image     *Image     `xml:"image,omitempty"`
	Mesh      *Mesh      `xml:"mesh,omitempty"`
	Plane     *Plane     `xml:"plane,omitempty"`
	Polyline  *Polyline  `xml:"polyline,omitempty"`
	Sphere    *Sphere    `xml:"sphere,omitempty"`
}

func (g *Geometry) Validate() error {
	gType := reflect.TypeOf(g).Elem()
	val := reflect.ValueOf(g).Elem()
	settedFieldName := ""
	for i := 0; i < val.NumField(); i++ {
		fType := gType.Field(i)
		if fType.Name == "XMLName" {
			continue
		}
		fValue := val.Field(i)
		if fValue.IsNil() == true {
			continue
		}
		if len(settedFieldName) > 0 {
			return fmt.Errorf("Multiple geometry '%s' in sdf.Geometry",
				strings.Join([]string{settedFieldName, fType.Name}, ","))
		}
		settedFieldName = fType.Name
	}

	return nil
}

func NewBox(size Vec3) *Geometry {
	return &Geometry{Box: &Box{Size: size}}
}

func NewCylinder(radius, length float64) *Geometry {
	return &Geometry{Cylinder: &Cylinder{Radius: radius, Length: length}}
}

func (b *Box) Inertia(mass float64) *Inertia {
	return &Inertia{
		Xx: mass / 12.0 * (b.Size[1]*b.Size[1] + b.Size[2]*b.Size[2]),
		Xy: 0,
		Xz: 0,
		Yy: mass / 12.0 * (b.Size[0]*b.Size[0] + b.Size[2]*b.Size[2]),
		Yz: 0,
		Zz: mass / 12.0 * (b.Size[0]*b.Size[0] + b.Size[1]*b.Size[1]),
	}
}

func (b *Box) Mass(density float64) float64 {
	return density * b.Size[0] * b.Size[1] * b.Size[2]
}

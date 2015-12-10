package pollyanna

import (
	"fmt"
	"regexp"
	"strings"
)

// Path is a single <path/> node
type Path struct {
	Fill    string `xml:"fill,attr"`
	RawPath string `xml:"d,attr"`
}

// ToPolygon converts a <path> node to a Polygon object
func (p Path) ToPolygon() Polygon {
	var polygon Polygon

	polygon.Fill = p.Fill

	re := regexp.MustCompile(`(?:\d+(?:\.?\d+?))`)
	coords := re.FindAllString(p.RawPath, -1)
	polygon.RawPoints = strings.Join(coords, ` `)

	return polygon
}

func (p Path) String() string {
	return fmt.Sprintf("Path: fill - %s, points - %s", p.Fill, p.RawPath)
}

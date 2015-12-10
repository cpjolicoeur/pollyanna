package pollyanna

import (
	"fmt"
	"regexp"
	"strings"
)

// Polygon is a single <polygon/> node
type Polygon struct {
	Fill      string `xml:"fill,attr"`
	RawPoints string `xml:"points,attr"`
}

func (p Polygon) FormattedCSSPolygonPoints() string {
	points := p.Points()
	return fmt.Sprintf("polygon(%s)", cssPolygonBuilder(points[0], points[1:len(points)]))
}

func (p Polygon) Points() [][]string {
	var points [][]string

	re := regexp.MustCompile(`(?:\d+(?:\.\d+)?[, ]\d+(?:\.\d+)?)`)
	coords := re.FindAllString(strings.Trim(p.RawPoints, ` `), -1)

	for _, coord := range coords {
		// Normalize the coord pair to be comma separated, then split on that comma
		points = append(points, strings.Split(strings.Replace(coord, ` `, `,`, 1), `,`))
	}

	return points
}

func (p Polygon) String() string {
	return fmt.Sprintf("Polygon: fill - %s, points - %s", p.Fill, p.RawPoints)
}

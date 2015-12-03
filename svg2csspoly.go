package svg2csspoly

import (
	"encoding/xml"
	"fmt"
)

// Svg is the top level <svg/> node
type Svg struct {
	XMLName  xml.Name  `xml:"svg"`
	Version  string    `xml:"version,attr"`
	Width    string    `xml:"width,attr"`
	Height   string    `xml:"height,attr"`
	ViewBox  string    `xml:"viewBox,attr"`
	Polygons []Polygon `xml:"polygon"`
}

// Polygon is a single <polygon/> node
type Polygon struct {
	Fill   string `xml:"fill,attr"`
	Points string `xml:"points,attr"`
}

// ParseSVG will parse the incoming SVG document bytes
func ParseSVG(bytes []byte) (Svg, error) {
	var svg Svg

	err := xml.Unmarshal(bytes, &svg)
	if err != nil {
		return svg, err
	}

	return svg, nil
}

func (s Svg) String() string {
	return fmt.Sprintf("SVG: version - %s, width - %s, height - %s, box - %s", s.Version, s.Width, s.Height, s.ViewBox)
}

func (p Polygon) String() string {
	return fmt.Sprintf("Polygon: fill - %s, points - %s", p.Fill, p.Points)
}

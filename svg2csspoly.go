package svg2csspoly

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
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
	Fill      string `xml:"fill,attr"`
	RawPoints string `xml:"points,attr"`
}

// Output represents the associated HTML Dom and CSS markup for the SVG image
type Output struct {
	HTML     string
	CSS      string
	CSSNodes []cssNode
}

type cssNode struct {
	selector string
	rules    []cssRule
}

type cssRule struct {
	property string
	value    string
}

// ParseSVG will parse the incoming SVG document bytes
func ParseSVG(bytes []byte) (Svg, error) {
	var svg Svg

	err := xml.Unmarshal(bytes, &svg)
	if err != nil {
		return svg, err
	}

	if 0 == len(svg.Polygons) {
		return svg, errors.New("No <polygon/> nodes were found in the SVG data.")
	}

	return svg, nil
}

// GenerateOutput converts the raw SVG data into HTML DOM nodes and associated CSS rules
func (s Svg) GenerateOutput() (Output, error) {
	var output Output
	return output, nil
}

func (s Svg) String() string {
	return fmt.Sprintf("SVG: version - %s, width - %s, height - %s, box - %s", s.Version, s.Width, s.Height, s.ViewBox)
}

func (p Polygon) String() string {
	return fmt.Sprintf("Polygon: fill - %s, points - %s", p.Fill, p.RawPoints)
}

func (p Polygon) Points() [][]string {
	var points [][]string
	pointsStr := strings.Split(strings.Trim(p.RawPoints, ` `), ` `)

	for _, pointStr := range pointsStr {
		points = append(points, strings.Split(pointStr, `,`))
	}

	return points
}

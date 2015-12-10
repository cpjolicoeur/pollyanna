package pollyanna

import (
	"encoding/xml"
	"fmt"
	"strings"
)

var baseDOM = `<div class="pollyanna"><div class="wrap">%s</div></div>`
var shardDOM = `<div class="shard-wrap"><div class="shard"></div></div>`

var baseCSS = `.pollyanna{position:absolute;width:100%;height:100%;top:0;left:0}.wrap{width:800;height:600px;top:5%;left:5%;position:absolute}.shard-wrap,.shard-wrap .shard,.shard-wrap .shard::before{width:100%;height:100%;position:absolute}.shard-wrap{z-index:2}.shard-wrap .shard{background-color:#fff}.shard-wrap .shard::before{content:"";background:rgba(255,255,255,0);top:0;left:0}`
var shardCSS = `.shard-wrap:nth-child(%d) .shard{-webkit-clip-path:%s;clip-path:%s;background-color:%s;}`

// Output represents the associated HTML Dom and CSS markup for the SVG image
type Output struct {
	HTML string
	CSS  string
}

// Svg is the top level <svg/> node
type Svg struct {
	XMLName  xml.Name  `xml:"svg"`
	Version  string    `xml:"version,attr"`
	Width    string    `xml:"width,attr"`
	Height   string    `xml:"height,attr"`
	ViewBox  string    `xml:"viewBox,attr"`
	Polygons []Polygon `xml:"polygon"`
}

// GenerateOutput converts the raw SVG data into HTML DOM nodes and associated CSS rules
func (s Svg) GenerateOutput() (Output, error) {
	var output Output
	output.HTML = fmt.Sprintf(baseDOM, strings.Repeat(shardDOM, len(s.Polygons)))
	output.CSS = baseCSS + s.cssShardChildren()
	return output, nil
}

func (s Svg) cssShardChildren() string {
	css := make([]string, len(s.Polygons))

	for i, p := range s.Polygons {
		cssPolyString := p.FormattedCSSPolygonPoints()
		css[i] = fmt.Sprintf(shardCSS, i+1, cssPolyString, cssPolyString, p.Fill)
	}

	return strings.Join(css, ``)
}

func cssPolygonBuilder(points []string, rest [][]string) string {
	if len(points) == 0 {
		return ``
	}

	if len(rest) == 0 {
		return cssFormatSinglePoint(points)
	}

	fmtString := "%s, %s"
	if len(rest) == 1 {
		return fmt.Sprintf(fmtString, cssFormatSinglePoint(points), cssFormatSinglePoint(rest[0]))
	}

	return fmt.Sprintf(fmtString, cssFormatSinglePoint(points), cssPolygonBuilder(rest[0], rest[1:len(rest)]))
}

func cssFormatSinglePoint(p []string) string {
	return fmt.Sprintf("%spx %spx", p[0], p[1])
}

func (s Svg) String() string {
	return fmt.Sprintf("SVG: version - %s, width - %s, height - %s, box - %s", s.Version, s.Width, s.Height, s.ViewBox)
}

package pollyanna

import (
	"encoding/xml"
	"errors"
)

// ParseSVG will parse the incoming SVG document bytes
func ParseSVG(bytes []byte) (Svg, error) {
	var n Node
	var svg Svg

	err := xml.Unmarshal(bytes, &n)
	if err != nil {
		return svg, err
	}

	if "svg" != n.XMLName.Local {
		return svg, errors.New("No root-level <svg> node found in the input data.")
	}

	svg.Polygons = n.BuildPolygons()
	if 0 == len(svg.Polygons) {
		return svg, errors.New("No <polygon/> or <path/> nodes were found in the SVG data.")
	}

	return svg, nil
}

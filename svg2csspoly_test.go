package svg2csspoly

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParsing(t *testing.T) {
	Convey("Invalid SVG input", t, func() {
		Convey("Handles no data", func() {
			svgData := []byte(``)

			_, err := ParseSVG(svgData)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "EOF")
		})

		Convey("Handles SVG files that containing non-polygon nodes", func() {
			svgData := []byte(`
			<svg viewBox='0 0 125 80' xmlns='http://www.w3.org/2000/svg'>
				<text y="75" font-size="100" font-family="serif"><![CDATA[10]]></text>
				</svg>
				`)

			svg, err := ParseSVG(svgData)
			So(svg.Polygons, ShouldHaveLength, 0)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, `No <polygon/> nodes were found in the SVG data.`)
		})

		Convey("Handles SVG files with no <polygon/> nodes", func() {
			svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
			<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
			</svg>`)

			svg, err := ParseSVG(svgData)
			So(svg.Polygons, ShouldHaveLength, 0)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Valid SVG input", t, func() {
		svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
		<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
		<polygon fill="#F3CD5E" points="366.4,7.6 432.7,5.6 430.6,67.7 401.1,71.1 "/>
		<polygon fill="#F1BD36" points="432.7,5.6 441.3,66.2 401.1,71.1 "/>
		<polygon fill="#B77E00" points="366.4,7.6 364.1,67.7 401.1,71.1 "/>
		</svg>`)

		Convey("Properly parses all <polygon/> nodes", func() {
			svg, err := ParseSVG(svgData)
			So(len(svg.Polygons), ShouldEqual, 3)
			So(err, ShouldBeNil)
		})

		Convey("Each Polygon point knows it's Fill", func() {
			svg, _ := ParseSVG(svgData)

			So(svg.Polygons[0].Fill, ShouldEqual, `#F3CD5E`)
			So(svg.Polygons[1].Fill, ShouldEqual, `#F1BD36`)
			So(svg.Polygons[2].Fill, ShouldEqual, `#B77E00`)
		})

		Convey("Each Polygon point knows it's Points", func() {
			svg, _ := ParseSVG(svgData)

			expectedPoints := [][]string{
				[]string{"366.4", "7.6"},
				[]string{"432.7", "5.6"},
				[]string{"430.6", "67.7"},
				[]string{"401.1", "71.1"}}
			So(svg.Polygons[0].Points(), ShouldHaveLength, 4)
			So(svg.Polygons[0].Points(), ShouldResemble, expectedPoints)
		})
	})
}

func TestGenerateOutput(t *testing.T) {
	Convey("Builds HTML DOM nodes", t, nil)
	Convey("Builds CSS rules", t, nil)
	Convey("Builds Raw CSS rule data", t, nil)
}
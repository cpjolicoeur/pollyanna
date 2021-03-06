package pollyanna

import (
	"strings"
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

		Convey("Handles non-svg xml data", func() {
			svgData := []byte(`<parent><child1>foo</child1><child2 /></parent>`)

			_, err := ParseSVG(svgData)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "No root-level <svg> node found in the input data.")
		})

		Convey("Handles SVG files that containing non-polygon or path nodes", func() {
			svgData := []byte(`
			<svg viewBox='0 0 125 80' xmlns='http://www.w3.org/2000/svg'>
				<text y="75" font-size="100" font-family="serif"><![CDATA[10]]></text>
				</svg>
				`)

			svg, err := ParseSVG(svgData)
			So(svg.Polygons, ShouldHaveLength, 0)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, `No <polygon/> or <path/> nodes were found in the SVG data.`)
		})

		Convey("Handles SVG files with no <polygon/> or <path/> nodes", func() {
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
		<polygon fill="#B77E00" points="366.4 7.6 364.1 67.7 401.1 71.1 "/>
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

		Convey("Handles SVG files with only <path> nodes", func() {
			svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
			<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
				<path fill="#F1BD36" d="M432.7,5.6 L441.3,66.2 L401.1,71.1 "/>
				<path fill="#F1BD36" d="M 432.7 5.6 L 441.3 66.2 L 401.1 71.1 "/>
				<path fill="#F1BD36" d="M 432.7 5.6 L 441.3 66.2 401.1 71.1 "/>
			</svg>`)

			svg, err := ParseSVG(svgData)
			So(svg.Polygons, ShouldHaveLength, 3)
			So(err, ShouldBeNil)
		})

		Convey("Handles SVG files with both <polygon> and <path> nodes", func() {
			svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
			<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
				<polygon fill="#F1BD36" points="432.7,5.6 441.3,66.2 401.1,71.1 "/>
				<path fill="#F1BD36" d="M 432.7 5.6 L 441.3 66.2 L 401.1 71.1 "/>
			</svg>`)

			svg, err := ParseSVG(svgData)
			So(svg.Polygons, ShouldHaveLength, 2)
			So(err, ShouldBeNil)
		})
	})
}

func TestGenerateOutput(t *testing.T) {
	Convey("Invalid SVG Data:", t, nil)

	Convey("Valid SVG Polygon Data:", t, func() {
		svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
		<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
			<polygon fill="#F3CD5E" points="366.4,7.6 432.7,5.6 430.6,67.7 401.1,71.1 "/>
			<polygon fill="#F1BD36" points="432.7,5.6 441.3,66.2 401.1,71.1 "/>
			<polygon fill="#B77E00" points="366.4 7.6 364.1 67.7 401.1 71.1 "/>
		</svg>`)
		svg, _ := ParseSVG(svgData)
		output, _ := svg.GenerateOutput()

		Convey("HTML DOM nodes", func() {
			Convey("Builds full output", func() {
				expected := `<div class="pollyanna"><div class="wrap"><div class="shard-wrap"><div class="shard"></div></div><div class="shard-wrap"><div class="shard"></div></div><div class="shard-wrap"><div class="shard"></div></div></div></div>`
				So(output.HTML, ShouldEqual, expected)
			})

			Convey("Builds one shard per Polygon", func() {
				So(strings.Count(output.HTML, "shard-wrap"), ShouldEqual, len(svg.Polygons))
			})
		})

		Convey("CSS rules", func() {
			Convey("Builds one nth-child shard-wrap rule per Polygon", func() {
				So(strings.Count(output.CSS, "nth-child"), ShouldEqual, len(svg.Polygons))
			})

			Convey("Builds the right color per Polygon", func() {
				for _, p := range svg.Polygons {
					So(output.CSS, ShouldContainSubstring, p.Fill)
				}
			})

			Convey("Builds the right points per Polygon", func() {
				So(output.CSS, ShouldContainSubstring, "polygon(366.4px 7.6px, 432.7px 5.6px, 430.6px 67.7px, 401.1px 71.1px)")
				So(output.CSS, ShouldContainSubstring, "polygon(432.7px 5.6px, 441.3px 66.2px, 401.1px 71.1px)")
				So(output.CSS, ShouldContainSubstring, "polygon(366.4px 7.6px, 364.1px 67.7px, 401.1px 71.1px)")
			})
		})
	})

	Convey("Valid SVG Path Data:", t, func() {
		svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
		<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
			<path fill="#F1BD36" d="M432.7,5.6 L441.3,66.2 L401.1,71.1 "/>
			<path fill="#F1BD36" d="M 32.7 15.6 L 44.3 6.2 L 1.1 71.1 "/>
			<path fill="#F1BD36" d="M 42.7 8.6 L 41.0 66 40.1 71.1 "/>
		</svg>`)
		svg, _ := ParseSVG(svgData)
		output, _ := svg.GenerateOutput()

		Convey("Converts all path nodes to Polygons", func() {
			So(len(svg.Polygons), ShouldEqual, 3)
		})

		Convey("Builds one shard per Polygon", func() {
			So(strings.Count(output.HTML, "shard-wrap"), ShouldEqual, len(svg.Polygons))
		})

		Convey("Builds the right points per Polygon", func() {
			So(output.CSS, ShouldContainSubstring, "polygon(432.7px 5.6px, 441.3px 66.2px, 401.1px 71.1px)")
			So(output.CSS, ShouldContainSubstring, "polygon(32.7px 15.6px, 44.3px 6.2px, 1.1px 71.1px)")
			So(output.CSS, ShouldContainSubstring, "polygon(42.7px 8.6px, 41.0px 66px, 40.1px 71.1px)")
		})
	})

	Convey("Valid SVG Polygon & Path Data:", t, func() {
		svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
		<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
			<polygon fill="#F3CD5E" points="366.4,7.6 432.7,5.6 430.6,67.7 401.1,71.1 "/>
			<path fill="#F1BD36" d="M 432.7 5.6 L 441.3 66.2 401.1 71.1 "/>
		</svg>`)
		svg, _ := ParseSVG(svgData)
		output, _ := svg.GenerateOutput()

		Convey("Converts all polygon and path nodes to Polygons", func() {
			So(len(svg.Polygons), ShouldEqual, 2)
		})

		Convey("Builds one shard per Polygon", func() {
			So(strings.Count(output.HTML, "shard-wrap"), ShouldEqual, len(svg.Polygons))
		})
	})
}

func TestStructDescription(t *testing.T) {
	Convey("Valid SVG File", t, func() {
		svgData := []byte(`<?xml version="1.0" encoding="utf-8"?>
		<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
		<polygon fill="#F3CD5E" points="366.4,7.6 432.7,5.6 430.6,67.7 401.1,71.1 "/>
		<polygon fill="#F1BD36" points="432.7,5.6 441.3,66.2 401.1,71.1 "/>
		<polygon fill="#B77E00" points="366.4,7.6 364.1,67.7 401.1,71.1 "/>
		</svg>`)
		svg, _ := ParseSVG(svgData)

		Convey("SVG String", func() {
			So(svg.String(), ShouldStartWith, "SVG:")
		})

		Convey("Polygon String", func() {
			So(svg.Polygons[0].String(), ShouldStartWith, "Polygon:")
		})
	})
}

func TestCssPolygonBuilder(t *testing.T) {
	Convey("Handles no Points", t, func() {
		out := cssPolygonBuilder([]string{}, [][]string{})
		So(out, ShouldEqual, ``)
	})

	Convey("Handles single Point", t, func() {
		out := cssPolygonBuilder([]string{"1", "2"}, [][]string{})
		So(out, ShouldEqual, `1px 2px`)
	})

	Convey("Handles multiple Points", t, func() {
		out := cssPolygonBuilder([]string{"1", "2"}, [][]string{[]string{"3", "4"}, []string{"5.1", "6.03"}})
		So(out, ShouldEqual, `1px 2px, 3px 4px, 5.1px 6.03px`)
	})
}

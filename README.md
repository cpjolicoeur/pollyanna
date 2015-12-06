![Pollyanna SVG](http://cpjolicoeur.s3.amazonaws.com/pollyanna.svg)

# Pollyanna

[![Build Status](https://travis-ci.org/cpjolicoeur/pollyanna.svg?branch=master)](https://travis-ci.org/cpjolicoeur/pollyanna)

Go library to convert SVG data to HTML + CSS Polygons

## Usage

```go
svgData := `<?xml version="1.0" encoding="utf-8"?>
<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="800px" height="600px" viewBox="0 0 800 600" enable-background="new 0 0 800 600" xml:space="preserve">
<polygon fill="#F3CD5E" points="366.4,7.6 432.7,5.6 430.6,67.7 401.1,71.1 "/>
<polygon fill="#F1BD36" points="432.7,5.6 441.3,66.2 401.1,71.1 "/>
<polygon fill="#B77E00" points="366.4,7.6 364.1,67.7 401.1,71.1 "/>
</svg>`

svg, err := pollyanna.ParseSVG(svgData)

output, err = svg.GenerateOutput()

// Access compacted generated data
fmt.Println("Generated HTML DOM:", output.HTML)
fmt.Println("Generated CSS:", output.CSS)

// Access generated CSS rules directed
fmt.Println("Raw CSS Nodes:", output.CSSNodes)
```

## Note

Current version only supports flattend (no layers) SVG files and `<polygon>` nodes only.  When you export from Sketch, Photoshop, Illustrator, etc... please flatten your image with no layers.

Support for `<path>` nodes coming next followed but layer manipulation and parsing.


## Development

#### Requirements

* [GoConvey][1] (for test filesystem)

Please add tests where applicable when submitting a Pull Request.


[1]:https://github.com/smartystreets/goconvey

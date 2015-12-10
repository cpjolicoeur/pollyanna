package pollyanna

import (
	"encoding/xml"
	"fmt"
)

// Node is a generic SVG XML node
type Node struct {
	XMLName   xml.Name
	Nodes     []Node `xml:",any"`
	Width     string `xml:"width,attr"`
	Height    string `xml:"height,attr"`
	ViewBox   string `xml:"viewBox,attr"`
	Fill      string `xml:"fill,attr"`
	RawPoints string `xml:"points,attr"`
	D         string `xml:"d,attr"`
	ID        string `xml:"id,attr"`
}

func (node Node) BuildPolygons() []Polygon {
	polygons := []Polygon{}

	walk([]Node{node}, func(n Node) bool {
		switch n.XMLName.Local {
		case "polygon":
			polygons = append(polygons, Polygon{n.Fill, n.RawPoints})
		case "path":
			p := Path{n.Fill, n.D}
			polygons = append(polygons, p.ToPolygon())
		}
		return true
	})

	return polygons
}

func (n Node) String() string {
	return fmt.Sprintf("Node: ID - %s, Fill - %s", n.ID, n.Fill)
}

func walk(nodes []Node, f func(Node) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}

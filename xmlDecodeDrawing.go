package excelize

import "encoding/xml"

// decodeTwoCellAnchor directly maps the twoCellAnchor (Two Cell Anchor Shape
// Size). This element specifies a two cell anchor placeholder for a group, a
// shape, or a drawing element. It moves with cells and its extents are in EMU
// units.
type decodeTwoCellAnchor struct {
	EditAs  string `xml:"editAs,attr,omitempty"`
	Content string `xml:",innerxml"`
}

// decodeWsDr directly maps the root element for a part of this content type
// shall wsDr. In order to solve the problem that the label structure is changed
// after serialization and deserialization, two different structures are
// defined. decodeWsDr just for deserialization.
type decodeWsDr struct {
	A             string                 `xml:"xmlns a,attr"`
	Xdr           string                 `xml:"xmlns xdr,attr"`
	TwoCellAnchor []*decodeTwoCellAnchor `xml:"twoCellAnchor,omitempty"`
	XMLName       xml.Name               `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing wsDr,omitempty"`
}

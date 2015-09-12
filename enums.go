package htmlparser

// AttrStatus indicate a status of an attribute
type AttrStatus uint8

const (
	ASValid AttrStatus = iota
	ASDeprecated
	ASUnknown
)

// Type of HTML Element according to the HTML 5.0 spec
type HtmlElementType uint8

const (
	HETPhrasing   HtmlElementType = 0x1  // former "inline element"
	HETFlow                       = 0x2  // former "block element"
	HETMeta                       = 0x4  // control elements
	HETText                       = 0x8  // text block
	HETNRCharData                 = 0x10 // Non-Replaceable Char Data

	HETAnyContent  = HETPhrasing | HETFlow | HETText
	HETTransparent = HETPhrasing | HETFlow
	HETNone        = 0
)

type HtmlTagFormatting uint8

const (
	HTFSingle          HtmlTagFormatting = iota // Has no closing tag, e.g. <br>
	HTFOptionalClosing                          // has an optional closing tag, e.g. <li>
	HTFComplete                                 // must have a closing tag
)

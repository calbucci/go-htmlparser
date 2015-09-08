package htmlparser

type AttrStatus uint8
const (
		ASValid AttrStatus = iota
		ASDeprecated
		ASUnknown
)

type HtmlElementType uint8
const (
		HEPhrasing HtmlElementType = 0x1 // former "inline element"
		HEFlow = 0x2 // former "block element"
		HEMeta = 0x4 // control elements
		HEText = 0x8 // text block
		HENRCharData = 0x10 // Non-Replaceable Char Data

		HEAnyContent = HEPhrasing | HEFlow | HEText
		HETransparent = HEPhrasing | HEFlow
		HENone = 0
)

type HtmlTagFormatting uint8
const (
		HTFSingle HtmlTagFormatting = iota // Has no closing tag, e.g. <br>
		HTFOptionalClosing // has an optional closing tag, e.g. <li>
		HTFComplete // must have a closing tag
)

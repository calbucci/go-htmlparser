package htmlparser

import (
	"bytes"
	"html"
	"strings"
	"unicode/utf8"
)

const maxInnerTextLengthStored = 65500

var ignoreTextInsideTag map[string]bool

type TextCallback func(string, *HtmlElement)
type ElementCallback func(*HtmlElement, bool)
type EndElementCallback func(string)

type HtmlParser struct {
	OrigHtml string
	stop     bool

	textCallback       TextCallback
	elementCallback    ElementCallback
	endElementCallback EndElementCallback

	Errors   []string
	Warnings []string

	Ids map[string]bool

	innerTextBuilder *bytes.Buffer
	InnerText string
	
	HasValidSyntax bool
	HasOnlyValidTags bool
	HasOnlyValidAttributes bool
	HasOnlyKnownTags bool
	HasDeprecatedAttributes bool
	HasDeprecatedTags bool
	
	SkipComments bool
	PreserveCRLFTab bool
	
	
}

func init() {
	ignoreTextInsideTag = map[string]bool{
		"head":   true,
		"html":   true,
		"ol":     true,
		"select": true,
		"table":  true,
		"tbody":  true,
		"thead":  true,
		"tfoot":  true,
		"tr":     true,
	}
}

func NewParser(html string, textCallback TextCallback, elementCallback ElementCallback, endElementCallback EndElementCallback) HtmlParser {
	var parser HtmlParser
	
	parser.OrigHtml = html
	parser.textCallback = textCallback
	parser.elementCallback = elementCallback
	parser.endElementCallback = endElementCallback
	
	parser.SkipComments = true
	parser.PreserveCRLFTab = true
	
	parser.innerTextBuilder = bytes.NewBufferString("")

	return parser	
}

func (parser *HtmlParser) Parse() bool {
	if(parser.stop){
		return false
	}
	
	parser.HasValidSyntax = true
			parser.HasOnlyValidTags = true
			parser.HasOnlyValidAttributes = true
			parser.HasDeprecatedAttributes = false
			parser.HasDeprecatedTags = false
			parser.HasOnlyKnownTags = true
			
	parser.Errors = make([]string,1)
	parser.Warnings = make([]string,1)
	parser.Ids = make(map[string]bool, 1)
	
	if(parser.OrigHtml == ""){
		return true
	}
	
	if(strings.Index(parser.OrigHtml, "<") < 0){
		parser.callText(parser.OrigHtml, nil)
	} else {
		parser.internalParse()	
	}
	
	parser.InnerText = html.UnescapeString(parser.innerTextBuilder.String())
	parser.stop = true
	
	return parser.HasValidSyntax
}

func (parser *HtmlParser) IsValidStrictHTML401() bool{
	return parser.HasValidSyntax && parser.HasOnlyValidTags && parser.HasOnlyValidAttributes
}

func (parser *HtmlParser) IsValidStrictHTMLNoDeprecated() bool{
	return parser.HasValidSyntax && parser.HasOnlyValidTags && parser.HasOnlyValidAttributes && !parser.HasDeprecatedAttributes && !parser.HasDeprecatedTags
}

func (parser *HtmlParser) IsValidHTML401() bool {
	return parser.HasValidSyntax && parser.HasOnlyValidTags && parser.HasOnlyValidAttributes
}

func (parser *HtmlParser) Stop() {
	parser.stop = true
}

func (p *HtmlParser) CallText(text string, parent *HtmlElement) {
	
	if(text == ""){
		return
	}
	
			if (!p.PreserveCRLFTab){
				if(!hasContent(text)){
					return
				}
			}

			if (parent != nil && parent.ElementInfo != nil){
				var childrenTypes = parent.ElementInfo.PermittedChildrenTypes;
				if ((childrenTypes & (HETText | HETNRCharData)) == 0){
					p.addWarning("Text node inside a " + parent.TagName + " element is not valid")
				}
			}

			if (parent != nil) {
				_,present := ignoreTextInsideTag[parent.TagNameNS]
				if(present){
					return
				}
			} 
			
			clearText := !p.PreserveCRLFTab

			if (parent != nil) {
				switch (parent.TagNameNS) {
					case "pre":
					clearText = false
					case "script":
					clearText = false
					case "style":
						clearText = false
				}
			}
			
			if (clearText && strings.HasPrefix(text, "<!--")){
					clearText = false
			}

			if (clearText) {
				// Remove all tags, CR, LF
				text = trimInBetween(text)
			}

			if (p.textCallback != nil) {
				p.textCallback(text, parent)
				if(p.stop) {
					return
					}
			}

			useBlock := false
			if (parent != nil && parent.ElementInfo != nil){
				if (parent.ElementInfo.ElementType == HETFlow) {
					useBlock = true
				} else {
					// This is whacky. We messed up on the definition of block-level elements for TR/TD
					// because of the optional tags.
					switch (parent.TagNameNS) {
						case "tr":
							useBlock = true
						case "td":
							useBlock = true
					}
				}
			}

			l := p.innerTextBuilder.Len()
			if (l < maxInnerTextLengthStored) {
				if (useBlock) {
					prevChar := '\n'
					if (l > 0) {
						prevChar,_ = utf8.DecodeLastRune(p.innerTextBuilder.Bytes())
					}

					if (prevChar != '\n' && prevChar != '\r') {
						p.innerTextBuilder.WriteRune('\n')
					}
					p.innerTextBuilder.WriteString(text)
					p.innerTextBuilder.WriteRune('\n')
				} else {
					p.innerTextBuilder.WriteString(text)
				}
			}
		}






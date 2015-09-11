package htmlparser

import (
	"bytes"
	"html"
	"strconv"
	"strings"
	"unicode/utf8"
	//"fmt"
)

const maxInnerTextLengthStored = 65500

var ignoreTextInsideTag map[string]bool

type TextCallback func(string, *HtmlElement)
type ElementCallback func(*HtmlElement, bool)
type EndElementCallback func(string)

type HtmlParser struct {
	OrigHtml  string
	origRunes []rune
	stop      bool

	textCallback       TextCallback
	elementCallback    ElementCallback
	endElementCallback EndElementCallback

	Errors   []string
	Warnings []string

	Ids map[string]bool

	innerTextBuilder *bytes.Buffer
	InnerText        string

	HasValidSyntax          bool
	HasOnlyValidTags        bool
	HasOnlyValidAttributes  bool
	HasOnlyKnownTags        bool
	HasDeprecatedAttributes bool
	HasDeprecatedTags       bool

	SkipComments    bool
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

func NewParser(html string) HtmlParser {
	var parser HtmlParser

	parser.OrigHtml = html

	parser.SkipComments = true
	parser.PreserveCRLFTab = true

	parser.innerTextBuilder = bytes.NewBufferString("")

	return parser
}

func (parser *HtmlParser) Parse(textCallback TextCallback, elementCallback ElementCallback, endElementCallback EndElementCallback) bool {
	if parser.stop {
		return false
	}

	parser.textCallback = textCallback
	parser.elementCallback = elementCallback
	parser.endElementCallback = endElementCallback

	parser.HasValidSyntax = true
	parser.HasOnlyValidTags = true
	parser.HasOnlyValidAttributes = true
	parser.HasDeprecatedAttributes = false
	parser.HasDeprecatedTags = false
	parser.HasOnlyKnownTags = true

	parser.Errors = make([]string, 1)
	parser.Warnings = make([]string, 1)
	parser.Ids = make(map[string]bool, 1)

	if parser.OrigHtml == "" {
		return true
	}

	if strings.Index(parser.OrigHtml, "<") < 0 {
		parser.callText(parser.OrigHtml, nil)
	} else {
		parser.internalParse()
	}

	parser.InnerText = html.UnescapeString(parser.innerTextBuilder.String())
	parser.stop = true

	return parser.HasValidSyntax
}

func (parser *HtmlParser) IsValidStrictHTML401() bool {
	return parser.HasValidSyntax && parser.HasOnlyValidTags && parser.HasOnlyValidAttributes
}

func (parser *HtmlParser) IsValidStrictHTMLNoDeprecated() bool {
	return parser.HasValidSyntax && parser.HasOnlyValidTags && parser.HasOnlyValidAttributes && !parser.HasDeprecatedAttributes && !parser.HasDeprecatedTags
}

func (parser *HtmlParser) IsValidHTML401() bool {
	return parser.HasValidSyntax && parser.HasOnlyValidTags && parser.HasOnlyValidAttributes
}

func (parser *HtmlParser) Stop() {
	parser.stop = true
}

func (p *HtmlParser) callText(text string, parent *HtmlElement) {

	if text == "" {
		return
	}

	if !p.PreserveCRLFTab {
		if !hasContent(text) {
			return
		}
	}

	if parent != nil && parent.ElementInfo != nil {
		var childrenTypes = parent.ElementInfo.PermittedChildrenTypes
		if (childrenTypes & (HETText | HETNRCharData)) == 0 {
			p.addWarning("Text node inside a " + parent.TagName + " element is not valid")
		}
	}

	if parent != nil {
		_, present := ignoreTextInsideTag[parent.TagNameNS]
		if present {
			return
		}
	}

	clearText := !p.PreserveCRLFTab

	if parent != nil {
		switch parent.TagNameNS {
		case "pre":
			clearText = false
		case "script":
			clearText = false
		case "style":
			clearText = false
		}
	}

	if clearText && strings.HasPrefix(text, "<!--") {
		clearText = false
	}

	if clearText {
		// Remove all tabs, CR, LF
		n := bytes.NewBufferString("")
		for _, r := range text {
			if r == '\n' {
				n.WriteRune(' ')
			} else if r >= rune(32) {
				n.WriteRune(r)
			}
		}
		text = n.String()
	}

	text = html.UnescapeString(text)

	if p.textCallback != nil {
		p.textCallback(text, parent)
		if p.stop {
			return
		}
	}

	useBlock := false
	if parent != nil && parent.ElementInfo != nil {
		if parent.ElementInfo.ElementType == HETFlow {
			useBlock = true
		} else {
			// This is whacky. We messed up on the definition of block-level elements for TR/TD
			// because of the optional tags.
			switch parent.TagNameNS {
			case "tr":
				useBlock = true
			case "td":
				useBlock = true
			}
		}
	}

	l := p.innerTextBuilder.Len()
	if l < maxInnerTextLengthStored {
		if useBlock {
			prevChar := '\n'
			if l > 0 {
				prevChar, _ = utf8.DecodeLastRune(p.innerTextBuilder.Bytes())
			}

			if prevChar != '\n' && prevChar != '\r' {
				p.innerTextBuilder.WriteRune('\n')
			}
			p.innerTextBuilder.WriteString(text)
			p.innerTextBuilder.WriteRune('\n')
		} else {
			p.innerTextBuilder.WriteString(text)
		}
	}
}

func (hp *HtmlParser) internalParse() {
	openedTags := make([]*HtmlElement, 0)
	openedBlocks := make([]*HtmlElement, 0)

	anyContent := false

	var parent, blockParent *HtmlElement

	fatal := false

	var text string
	last := 0

	var c rune
	hp.origRunes = []rune(hp.OrigHtml)
	l := len(hp.origRunes)
	//fmt.Printf("Len: %v\n", l)
	len1 := l - 1

	for p := 0; p < l; p++ {
		//fmt.Printf("p=%v | last=%v\n", p, last)
		if hp.stop {
			return
		}
		c = hp.origRunes[p]

		if c != '<' {
			continue
		}

		diff := p - last
		if diff >= 1 {
			parent = nil
			if hp.HasValidSyntax && len(openedTags) > 0 {
				parent = openedTags[len(openedTags)-1]
			}
			//fmt.Printf("1-[%v:%v]\n", last, diff + last)
			text2 := string(hp.origRunes[last : diff+last])
			if hasContent(text2) {
				anyContent = true
			}
			hp.callText(text2, parent)
			if hp.stop {
				return
			}
		}
		// Yes, this is an open (or closing) tag
		if p >= len1 {
			hp.HasValidSyntax = false
			fatal = true
			hp.addError("HTML ends with < character")
			break // the html ends with this open tag
		}

		startElem := p
		elem := hp.getElementString(startElem)
		if elem == "" {
			fatal = true
			break // fatal syntax error
		}
		ecl := utf8.RuneCountInString(elem)
		//fmt.Printf("ecl=%v | elem='%v'\n", ecl, elem)
		p += ecl - 1
		if ecl <= 2 {
			// bad HTML, like "<>"
			hp.addError("Element is empty <>")
			hp.HasValidSyntax = false
			continue
		}
		last = p + 1
		parent, blockParent = nil, nil
		if hp.HasValidSyntax {
			if len(openedTags) > 0 {
				parent = openedTags[len(openedTags)-1]
			}
			if len(openedBlocks) > 0 {
				blockParent = openedBlocks[len(openedBlocks)-1]
			}
		}

		if elem[1] == '/' {
			anyContent = true
			// This is a closing tag
			tag := parseClosingTag(elem)
			//fmt.Printf("309tag: %v\n", tag)

			if hp.HasValidSyntax {
				hp.unwindForClose(tag, &openedTags, &openedBlocks)
				//fmt.Printf("313-openTags: %v | openedBlocks: %v\n", openedTags, openedBlocks)
				if hp.stop {
					return
				}
			}
			if hp.endElementCallback != nil {
				hp.endElementCallback(tag)
				if hp.stop {
					return
				}
			}
		} else {
			if strings.HasPrefix(elem, "<!") {
				if strings.HasPrefix(elem, "<!--") {
					p = p - utf8.RuneCountInString(elem) + 1
					// It's a comment
					ec := runesIndexRunesStart(hp.origRunes, []rune("-->"), p+4)
					if ec == -1 {
						hp.HasValidSyntax = false
						fatal = true
						hp.addError("Missing end comment -->")
						break
					}
					//fmt.Printf("2-[%v:%v]\n", p, ec + 3)
					text = string(hp.origRunes[p : ec+3])
					if !hp.SkipComments {
						hp.callText(text, nil)
						if hp.stop {
							return
						}
					}
					p += utf8.RuneCountInString(text)
					last = p
					p--
					continue
				}
				// Looks like a doctype
				e2 := strings.ToLower(elem)
				//fmt.Printf("351-e2=%v\n", e2)
				if strings.HasPrefix(e2, "<!doctype ") {
					if anyContent {
						hp.HasValidSyntax = false
						hp.addError("The doctype declaration must be the first element of the HTML:" + e2)
					}
					anyContent = true
					hp.callText(elem, nil)
					if hp.stop {
						return
					}
					continue
				}
			}
			anyContent = true
			he := NewHtmlElement(elem, parent, &hp.Errors, &hp.Warnings)
			if he.HasDeprecatedAttributes {
				hp.HasDeprecatedAttributes = true
			}
			if !he.HasOnlyKnownAttributes {
				hp.HasOnlyValidAttributes = false
			}

			if he.Id != "" {
				_, present := hp.Ids[he.Id]
				if present {
					hp.addWarning("Duplicate id: " + he.Id + " - Element: " + he.OriginalOpenTag)
				}
				hp.Ids[he.Id] = true
			}

			if he.SyntaxError {
				hp.HasValidSyntax = false
				if he.FatalSyntaxError {
					fatal = true
					break
				}
				hp.addError("Element syntax error for " + he.OriginalOpenTag)
				continue
			}

			// Special cases for script and style
			if he.TagNameNS == "script" || he.TagNameNS == "style" {
				startScript := p + 1
				endScript := 0
				endTag := ""
				cp := startScript
				for {
					cp = runesIndexRunesStart(hp.origRunes, []rune("</"), cp)
					if cp == -1 {
						hp.HasValidSyntax = false
						fatal = true
						hp.addError("Missing close tag: " + he.OriginalOpenTag)
						break
					}
					endScript = cp
					elem = hp.getElementString(cp)
					if elem == "" {
						hp.HasValidSyntax = false
						fatal = true
						hp.addError("Missing close > for closing tag: " + he.OriginalOpenTag)
						break
					}
					cp += utf8.RuneCountInString(elem)
					last = cp
					endTag = parseClosingTag(elem)
					if endTag == he.TagNameNS {
						break
					}
					endTag = ""
				}
				if endTag == "" {
					p = l
					break
				}
				p = cp - 1

				if hp.elementCallback != nil {
					hp.elementCallback(he, false)
					if hp.stop {
						return
					}
				}

				//fmt.Printf("3-[%v:%v]\n", startScript, endScript)
				hp.callText(string(hp.origRunes[startScript:endScript]), he)
				if hp.stop {
					return
				}

				if hp.endElementCallback != nil {
					hp.endElementCallback(he.TagNameNS)
					if hp.stop {
						return
					}
				}
				continue
			}

			// We consider this a single element if
			// 1) the ElementInfo.Single is flagged
			// 2) It is an unknown element (but not in any namespace)
			if he.ElementInfo == nil {
				hp.HasOnlyKnownTags = false
				// Unknown HTML 4.01 tag
				if !he.HasNamespace {
					hp.addWarning("Unknown tag: " + he.TagNameNS)
					// Really unknown and invalid tag
					if he.XmlEmptyTag {
						if hp.elementCallback != nil {
							hp.elementCallback(he, true)
							if hp.stop {
								return
							}
						}
					} else {
						if hp.elementCallback != nil {
							hp.elementCallback(he, false)
							if hp.stop {
								return
							}
						}
						if hp.HasValidSyntax {
							openedTags = append(openedTags, he)
						}
					}
				} else {
					// it is unknown, but correctly declared in an XML namespace
					if he.XmlEmptyTag {
						if hp.elementCallback != nil {
							hp.elementCallback(he, true)
							if hp.stop {
								return
							}
						}
					} else {
						if hp.elementCallback != nil {
							hp.elementCallback(he, false)
							if hp.stop {
								return
							}
						}
						if hp.HasValidSyntax {
							openedTags = append(openedTags, he)
						}
					}
				}
			} else {
				if he.ElementInfo.Obsolete {
					hp.addWarning("Deprecated Tag: " + he.TagNameNS)
					hp.HasDeprecatedTags = true
				}

				// It's  known tag
				if he.ElementInfo.TagFormatting == HTFSingle || he.XmlEmptyTag {
					if hp.elementCallback != nil {
						hp.elementCallback(he, true)
						if hp.stop {
							return
						}
					}
				} else {
					if hp.HasValidSyntax {
						if he.ElementInfo.ElementType == HETFlow {
							// Some Tags have optional closing (like LI or TD or P)
							// We assume an automatic closing for these tags on the following situation:
							// 1) Current element is block-level, and
							// 2) Parent node is also a block-level and supports optional closing
							// 3) Current element is the same class as parent element
							//    or Current element is the closing tag of the parent element
							if blockParent != nil && blockParent.ElementInfo.TagFormatting == HTFOptionalClosing {
								if parent != blockParent {
									hp.addWarning("Invalid parent for " + blockParent.TagName + " (inside of " + parent.TagName + ")")
								} else {
									if he.TagName == blockParent.TagName {
										if hp.endElementCallback != nil {
											hp.endElementCallback(parent.TagNameNS)
											if hp.stop {
												return
											}
										}
										openedTags = openedTags[:len(openedTags)-1]
										openedBlocks = openedBlocks[:len(openedBlocks)-1]
									}
								}
							}
							if hp.HasValidSyntax {
								openedBlocks = append(openedBlocks, he)
							}
						}
						if hp.HasValidSyntax {
							openedTags = append(openedTags, he)
						}
					}
					if hp.elementCallback != nil {
						hp.elementCallback(he, false)
					}

				}
			}

		}
	} // for loop

	//fmt.Printf("554-Out\n")

	if !fatal {
		// commit the last piece of text
		parent = nil
		//fmt.Printf("559\n")
		if hp.HasValidSyntax && len(openedTags) > 0 {
			parent = openedTags[len(openedTags)-1]
		}

		//fmt.Printf("564\n")
		if last < l {
			//fmt.Printf("564-[%v:]\n", last)
			text = string(hp.origRunes[last:])
			hp.callText(text, parent)
			if hp.stop {
				return
			}
		}

		//fmt.Printf("574\n")
		if hp.HasValidSyntax {
			//fmt.Printf("576-openedTags: %v %v\n", len(openedTags), openedTags)
			for len(openedTags) > 0 {
				parent = openedTags[len(openedTags)-1]
				if parent.ElementInfo == nil || parent.ElementInfo.TagFormatting != HTFOptionalClosing {
					break
				}
				if hp.endElementCallback != nil {
					hp.endElementCallback(parent.TagNameNS)
					if hp.stop {
						return
					}
				}
				openedTags = openedTags[:len(openedTags)-1]
				blockParent = nil
				if len(openedBlocks) > 0 {
					blockParent = openedBlocks[len(openedBlocks)-1]
				}
				if parent == blockParent {
					openedBlocks = openedBlocks[:len(openedBlocks)-1]
				}
			}
		}
		//fmt.Printf("598\n")
	}

	//fmt.Printf("596-OUt\n")

	if hp.HasValidSyntax {
		if len(openedBlocks) > 0 {
			if len(openedTags) != len(openedBlocks) {
				hp.HasValidSyntax = false
				hp.addError("Missing " + strconv.Itoa(len(openedTags)) + " tag(s) closing.")
			} else {
				for len(openedBlocks) > 0 {
					blockParent = openedBlocks[len(openedBlocks)-1]
					openedBlocks = openedBlocks[:len(openedBlocks)-1]
					parent = openedTags[len(openedTags)-1]
					openedTags = openedTags[:len(openedTags)-1]
					if parent != blockParent {
						hp.HasValidSyntax = false
						hp.addError("Missing a close tag for a block-element. Opened Tag: " + parent.TagNameNS)
						break
					}
					if hp.endElementCallback != nil {
						hp.endElementCallback(parent.TagNameNS)
						if hp.stop {
							return
						}
					}
				}

			}
		} else if len(openedTags) > 0 {
			hp.addError("Missing " + strconv.Itoa(len(openedTags)) + " tag(s) closing.")
			hp.HasValidSyntax = false
		}
	}
}

func (hp *HtmlParser) unwindForClose(tag string, openedTags, openedBlocks *[]*HtmlElement) {
	var parent, blockParent *HtmlElement
	if len(*openedTags) > 0 {
		parent = (*openedTags)[len(*openedTags)-1]
	}

	if parent == nil {
		hp.HasValidSyntax = false
		hp.addError("Closing tag without opening: " + tag)
		return
	}
	//fmt.Printf("637-Parent:%v\n", parent)

	firstParent := parent.TagNameNS

	if len(*openedBlocks) > 0 {
		blockParent = (*openedBlocks)[len(*openedBlocks)-1]
	}

	//fmt.Printf("645-openTags: %v | openedBlocks: %v\n", *openedTags, *openedBlocks)

	for parent != nil {
		if parent.TagNameNS == tag {
			*openedTags = (*openedTags)[:len(*openedTags)-1]
			//fmt.Printf("648-openTags: %v | openedBlocks: %v\n", *openedTags, *openedBlocks)
			if blockParent != nil && blockParent.TagNameNS == tag {
				*openedBlocks = (*openedBlocks)[:len(*openedBlocks)-1]
			}
			return
		}

		if parent.ElementInfo == nil {
			break // mismatch
		}

		// This could be either a tag mismatch, or an optional element missing
		if parent.ElementInfo.TagFormatting != HTFOptionalClosing {
			break // mismatch
		}

		// inject the optional closing tag
		if hp.endElementCallback != nil {
			hp.endElementCallback(parent.TagNameNS)
			if hp.stop {
				return
			}
		}

		if len(*openedTags) == 0 {
			break
		}
		*openedTags = (*openedTags)[:len(*openedTags)-1]
		if blockParent == parent {
			*openedBlocks = (*openedBlocks)[:len(*openedBlocks)-1]
			blockParent = nil
			if len(*openedBlocks) > 0 {
				blockParent = (*openedBlocks)[len(*openedBlocks)-1]
			}
		}
		parent = parent.Parent
	}

	hp.addError("Tag mismatch. Open tag: " + firstParent + " / Closing tag: " + tag)
	hp.HasValidSyntax = false
}

func (hp *HtmlParser) addError(error string) {
	hp.Errors = append(hp.Errors, error)
}

func (hp *HtmlParser) addWarning(wrn string) {
	hp.Warnings = append(hp.Warnings, wrn)
}

func (hp *HtmlParser) getElementString(startPos int) string {
	var c rune
	endElem := 0
	l := len(hp.origRunes)
	p := startPos
	for ; p < l; p++ {
		c = hp.origRunes[p]
		if c == '>' {
			endElem = p
			break
		}
		if c == '"' || c == '\'' {
			p = runesIndexRunesStart(hp.origRunes, []rune{c}, p+1)
			if p == -1 {
				// Not well formed HTML: <div attr="value>   (missing quote)
				logString := hp.origRunes[startPos:]

				if len(logString) > 40 {
					logString = logString[:40]
				}
				hp.addError("Attribute start quote but doesn't end with quote: " + string(logString))
				hp.HasValidSyntax = false
				return ""
			}
		}
	}

	if endElem == 0 {
		hp.HasValidSyntax = false
		logString := hp.origRunes[startPos:]
		if len(logString) > 40 {
			logString = logString[0:40]
		}
		hp.addError("Can't find > for tag: " + string(logString))
		return ""
	}

	return string(hp.origRunes[startPos : endElem+1])

}

package htmlparser

import (
	"bytes"
	"html"
	"strings"
	"unicode"
)

type QuoteType uint8

const (
	QTNone QuoteType = iota
	QTSingle
	QTDouble
)

type attributeInfo struct {
	Name  string
	Value string
}

type HtmlElement struct {
	errors   *[]string
	warnings *[]string

	TagName                 string
	TagNameNS               string
	Id                      string
	Attributes              []attributeInfo
	ElementInfo             *HtmlElementInfo
	Namespace               string
	HasNamespace            bool
	XmlEmptyTag             bool
	Parent                  *HtmlElement
	HasDeprecatedAttributes bool
	HasOnlyKnownAttributes  bool
	SyntaxError             bool
	FatalSyntaxError        bool
	OriginalOpenTag         string
}

func NewHtmlElement(openElement string, parent *HtmlElement, errors, warnings *[]string) *HtmlElement {

	he := new(HtmlElement)
	he.OriginalOpenTag = openElement

	he.Parent = parent

	he.errors = errors
	he.warnings = warnings

	he.HasOnlyKnownAttributes = true
	he.HasDeprecatedAttributes = false

	// openElement contains any type of open tag/single tag
	// Examples:
	//	<br>
	//  <br/>
	//	<br clear=left>
	//	<div style="color:#fff">
	//	<img src='/a/b/c'>

	he.Attributes = make([]attributeInfo, 0)

	runes := []rune(openElement)
	l := len(runes)

	pos := 1 // skip the <
	for ; pos < l; pos++ {
		c := runes[pos]
		if !unicode.IsSpace(c) {
			break
		}
	}

	if pos == l {
		// Error: Empty tag with whitespaces only: "<   >";
		he.addError("Invalid tag (whitespaces only).")
		he.SyntaxError = true
		return he
	}

	for ; pos < l; pos++ {
		c := runes[pos]
		if c == '>' {
			if pos == 1 {
				// Error: Empty tag like "<>"
				he.addError("Empty tag <>")
				he.SyntaxError = true
				return he
			}
			// This is it
			he.TagName = strings.ToLower(strings.TrimSpace(string(runes[1:pos])))
			he.checkTag()
			return he
		}

		if unicode.IsSpace(c) {
			he.TagName = strings.ToLower(strings.TrimSpace(string(runes[1:pos])))
			he.checkTag()
			break
		}
	}

	pos++ // skip the whitespace

	end := runesLastIndex(runes, '>')
	if end == -1 {
		he.addError("Missing closing >")
		he.SyntaxError = true
		he.FatalSyntaxError = true
		return he
	}
	end--
	for end >= pos {
		if runes[end] == '/' {
			he.XmlEmptyTag = true
			end--
			break
		}
		if !unicode.IsSpace(runes[end]) {
			break
		}
		end--
	}

	if end > pos {
		he.parseAttributes(squeezeSpaces(string(runes[pos : end+1])))
	}

	return he
}

func (he *HtmlElement) GetOpenTag(noEvents, noUnknownAttributes bool) string {
	return internalBuildOpenTag(he.ElementInfo, he.TagNameNS, he.Attributes, noEvents, noUnknownAttributes, he.XmlEmptyTag)
}

func (he *HtmlElement) GetCloseTag() string {
	return "</" + he.TagNameNS + ">"
}

func (he *HtmlElement) GetAttributeValue(attrName string) (string, bool) {

	i := he.FindAttributeIndex(attrName)
	if i >= 0 {
		return he.Attributes[i].Value, true
	}
	return "", false

}

func (he *HtmlElement) SetAttribute(attrName, attrValue string) bool {
	if attrName == "" {
		return true
	}

	if strings.IndexAny(attrValue, "\r\n\t") >= 0 {
		//throw new ArgumentException("attrValue cannot contain control characters")
		return false
	}

	i := he.FindAttributeIndex(attrName)
	if i >= 0 {
		he.Attributes[i].Value = attrValue
	} else {
		he.AddAttribute(attrName, attrValue)
	}
	return true
}

func (he *HtmlElement) RemoveAttribute(attrName string) {
	i := he.FindAttributeIndex(attrName)
	if i >= 0 {
		he.Attributes = append(he.Attributes[:i], he.Attributes[i+1:]...)
	}
}

func (he *HtmlElement) HasAttribute(attrName string) bool {
	return he.FindAttributeIndex(attrName) >= 0
}

func (he *HtmlElement) FindAttributeIndex(attrName string) int {
	if len(he.Attributes) == 0 || attrName == "" {
		return -1
	}

	attrName = strings.ToLower(attrName)

	for i, a := range he.Attributes {
		if a.Name == attrName {
			return i
		}
	}
	return -1
}

func (he *HtmlElement) checkTag() {
	if strings.HasSuffix(he.TagName, "/") {
		he.TagName = he.TagName[0 : len(he.TagName)-1]
	}
	he.TagNameNS = he.TagName

	he.ElementInfo = GetElementInfo(he.TagNameNS)

	pos := strings.Index(he.TagName, ":")
	if pos != -1 {
		he.Namespace = he.TagName[:pos]
		he.TagName = he.TagName[pos+1:]
	}
	if he.ElementInfo == nil {
		if he.Namespace == "" {
			he.addWarning("Unknown tag: " + he.TagName)
		}
	} else {
		if he.Parent != nil {
			if !he.ElementInfo.IsValidParent(he.Parent.TagName) {
				he.addWarning("Invalid parent for " + he.TagName + " (parent: " + he.Parent.TagName + ")")
			}
		}
	}
}

func (he *HtmlElement) addWarning(warning string) {
	*he.warnings = append(*he.warnings, warning)
}

func (he *HtmlElement) addError(error string) {
	*he.errors = append(*he.errors, error)
}

func (he *HtmlElement) AddAttribute(attrName, attrVal string) {
	if attrName == "" {
		return
	}

	if attrName == "style" {
		attrVal = cleanStyleAttr(attrVal)
	} else if attrName == "id" {
		he.Id = attrVal
	}

	if he.ElementInfo != nil {
		//bool useUrl;
		ast := he.ElementInfo.GetAttributeStatus(attrName)
		if ast == ASUnknown {
			if strings.Index(attrName, ":") > 0 {
			} else {
				he.HasOnlyKnownAttributes = false
				he.addWarning("Unknown attribute: " + attrName + " (tag: " + he.TagNameNS + ")")
			}
		} else if ast == ASDeprecated {
			he.HasDeprecatedAttributes = true
			he.addWarning("Deprecated attribute: " + attrName + " (tag: " + he.TagNameNS + ")")
		}
	}

	if len(attrVal) > 0 {
		attrVal = html.UnescapeString(attrVal)
	}

	he.Attributes = append(he.Attributes, attributeInfo{attrName, attrVal})
}

func squeezeSpaces(s string) string {
	n := bytes.NewBufferString("")
	atSpace := false
	atEqual := false
	inQuote := false
	quote := rune('-')

	for _, c := range s {

		if inQuote {
			if c == quote {
				inQuote = false
			}
			n.WriteRune(c)
			continue
		}
		if unicode.IsSpace(c) {
			atSpace = true
			continue
		}
		if c == '=' {
			atEqual = true
			continue
		}
		// At this point, we know the char is not white or '='.
		if atEqual {
			n.WriteRune('=')
			atEqual = false
			atSpace = false
		}
		if atSpace {
			n.WriteRune(' ')
			atSpace = false
		}
		if c == '"' || c == '\'' {
			inQuote = true
			quote = c
		}
		n.WriteRune(c)
	}

	if atEqual {
		n.WriteRune('=')
	}
	return n.String()
}

func (he *HtmlElement) parseAttributes(openElement string) {
	runes := []rune(openElement)
	l := len(runes)
	var attrName, attrVal string
	p := 0
	var c rune
	var found bool
	// Parse all the attributes now
	for ; p < l; p++ {
		// skip all the whitespaces
		for unicode.IsSpace(runes[p]) {
			p++
			if p == l {
				return
			}
		}

		// now, search for the attribute name by either finding a whitespace or the "=" sign
		found = false
		startAttrName := p
		for {
			c = runes[p]
			if unicode.IsSpace(c) || c == '>' {
				// This is an empty attribute like "checked" in "<input type=checkbox checked>"
				attrName = strings.ToLower(strings.TrimSpace(string(runes[startAttrName:p])))
				he.AddAttribute(attrName, "")
				if c == '>' {
					return
				}
				found = true
				break
			}
			if c == '=' {
				break
			}
			p++
			if p >= l {

				attrName = strings.ToLower(strings.TrimSpace(string(runes[startAttrName:p])))
				he.AddAttribute(attrName, "")
				return
			}
		}
		if found {
			continue
		}

		if startAttrName == p {
			he.addError("Attribute name starts with the '=' sign.")
			he.SyntaxError = true
			// Invalid attribute, starts with an '=' sign
			// Skip it to the next whitespace
			p++
			c = runes[p]
			if c == '\'' {
				p = p + 1 + runesIndex(runes[p+1:], '\'')
			} else if c == '"' {

				p = p + 1 + runesIndex(runes[p+1:], '"')
			}
			continue
		}

		attrName = strings.ToLower(strings.TrimSpace(string(runes[startAttrName:p])))
		p++ // skipt the equal sign
		if p == l {
			he.addError("Attribute ends with equal sign.")
			he.SyntaxError = true
			he.FatalSyntaxError = true
			return
		}

		startAttrVal := p
		c = runes[p]

		if unicode.IsSpace(c) || c == '>' {
			// This is a malformed attribute since it has a whitespace after the '=' sign,
			// like <a class= abc> or <a class=>
			he.addError("Attribute is missing value: " + attrName)
			he.SyntaxError = true
			he.AddAttribute(attrName, "")
			continue
		}

		if c == '\'' || c == '"' {
			startAttrVal++
			np := runesIndex(runes[p+1:], c)
			if np == -1 {
				// Argh, this attribute is missing the end quote, stop parsing
				he.addError("Attribute is missing end quote: " + attrName)
				he.SyntaxError = true
				he.FatalSyntaxError = true
				return
			}
			p = np + p + 1

			if p == startAttrVal {
				attrVal = ""
			} else {
				attrVal = string(runes[startAttrVal:p])
			}
			he.AddAttribute(attrName, attrVal)
			continue
		}

		// This is an attribute without a quote. Find the first whitespace or >
		for ; p < l; p++ {
			c = runes[p]
			if unicode.IsSpace(c) || c == '>' || p == l-1 {

				attrVal = string(runes[startAttrVal : p+1])
				he.AddAttribute(attrName, attrVal)
				break
			}
		}
	}

}

func parseClosingTag(elem string) string {
	if !strings.HasPrefix(elem, "</") {
		return ""
	}

	for p, c := range elem {

		if c == '>' || unicode.IsSpace(c) {

			return strings.ToLower(strings.TrimSpace(elem[2:p]))
		}
	}
	return strings.ToLower(strings.TrimSpace(elem))
}

func BuildOpenTagHEI(ei *HtmlElementInfo, attributes []attributeInfo, noEvents, noUnknownAttributes bool) string {
	return internalBuildOpenTag(ei, ei.TagName, attributes, noEvents, noUnknownAttributes, false)
}

func BuildOpenTag(tagName string, attributes []attributeInfo, noEvents, noUnknownAttributes bool) string {
	var ei *HtmlElementInfo
	if noUnknownAttributes {
		ei = GetElementInfo(tagName)
	}
	return internalBuildOpenTag(ei, tagName, attributes, noEvents, noUnknownAttributes, false)
}

func HtmlAttributeEncode(attributeValue string) string {

	if attributeValue == "" {
		return ""
	}

	if strings.IndexAny(attributeValue, `&"`) == -1 {
		return attributeValue
	}

	n := bytes.NewBufferString("")
	for _, c := range attributeValue {
		switch c {
		case '&':
			n.WriteString("&amp;")
		case '"':
			n.WriteString("&quot;")
		default:
			n.WriteRune(c)

		}
	}

	return n.String()
}

func NeedQuotesForAttr(val string) QuoteType {
	if val == "" {
		return QTDouble
	}

	qt := QTNone
	runes := []rune(val)
	for c := range runes {
		switch {
		case c >= 'a' && c <= 'z':
			continue
		case c >= 'A' && c <= 'Z':
			continue
		case c >= '0' && c <= '9':
			continue
		case c == '_' || c == '-' || c == '.' || c == ',': // According to http://www.w3.org/TR/html401/intro/sgmltut.html#h-3.2.2
			continue
		}
		qt = QTDouble
		if c == '"' {
			qt = QTSingle
		}
	}
	return qt
}

func cleanStyleAttr(style string) string {
	if style == "" {
		return style
	}

	parts := convertSemicolonDelimited(style)

	n := bytes.NewBufferString("")

	for _, part := range parts {
		p2 := strings.TrimSpace(part)
		if len(p2) == 0 {
			continue
		}
		pos := strings.IndexRune(p2, ':')
		if pos == -1 {
			continue
		}
		styleName := strings.ToLower(p2[:pos])
		styleValue := strings.TrimSpace(p2[pos+1:])

		if len(styleValue) == 0 {
			continue
		}

		if n.Len() > 0 {
			n.WriteRune(';')
		}
		n.WriteString(styleName)
		n.WriteRune(':')
		n.WriteString(styleValue)
	}

	return n.String()

}

func internalBuildOpenTag(ei *HtmlElementInfo, tagName string, attributes []attributeInfo, noEvents, noUnknownAttributes, xmlEmptyTag bool) string {
	if !noUnknownAttributes {
		ei = nil
	}

	n := bytes.NewBufferString("")

	n.WriteRune('<')
	n.WriteString(tagName)

	for _, a := range attributes {

		if a.Name == "" || noEvents && strings.HasPrefix(a.Name, "on") {
			continue
		}

		if ei != nil && ei.GetAttributeStatus(a.Name) == ASUnknown {
			continue
		}

		n.WriteRune(' ')
		n.WriteString(a.Name)
		if a.Value == "" {

			continue // Empty attribute (valid on HTML5 and above)
		}

		n.WriteRune('=')

		if len(a.Value) > 0 {
			encoded := html.EscapeString(a.Value)
			n.WriteRune('"')
			n.WriteString(encoded)
			n.WriteRune('"')
		}
	}
	if xmlEmptyTag {
		n.WriteString(" />")
	} else {
		n.WriteRune('>')
	}
	return n.String()

}

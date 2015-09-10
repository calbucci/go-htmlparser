package htmlparser

import (
	"strings"
)

var globalAttributes []string
var allElements []HtmlElementInfo
var elemsInfo map[string]HtmlElementInfo

func init() {
	baseAttributes := "accesskey;class;contenteditable;contextmenu;dir;draggable;dropzone;hidden;id;lang;spellcheck;style;tabindex;title;translate;;onabort;onblur;oncanplay;oncanplaythrough;onchange;onclick;oncontextmenu;ondblclick;ondrag;ondragend;ondragenter;ondragleave;ondragover;ondragstart;ondrop;ondurationchange;onemptied;onended;onerror;onfocus;oninput;oninvalid;onkeydown;onkeypress;onkeyup;onload;onloaddata;onloadeddata;onloadedmetadata;onloadstart;onmousedown;onmousemove;onmouseout;onmouseover;onmouseup;onmousewheel;onpause;onplay;onplaying;onprogress;onratechange;onreadystatechange;onreset;onscroll;onseekend;onseeking;onselect;onshow;onstalled;onsubmit;onsuspended;ontimeupdate;onvolumechange;onwaiting;xml:base;xml:lang;xml:space"
	globalAttributes = convertSemicolonDelimited(baseAttributes)

	initElements()

	elemsInfo = make(map[string]HtmlElementInfo, len(allElements))
	for _, hei := range allElements {
		hei.setAttributes(hei.attributesString)
		elemsInfo[hei.TagName] = hei
	}
}

type HtmlElementInfo struct {
	TagName                string
	HtmlVersion            int  // HTML version that introduced this tag
	Obsolete               bool // Indicates if this element is obsolete
	TagFormatting          HtmlTagFormatting
	ElementType            HtmlElementType
	PermittedChildrenTypes HtmlElementType // Valid types of elements that can be nested inside this tag
	PermittedChildrenTags  []string        // Valid children for this tag
	Attributes             []string
	attributesString       []string // This is temporary to be merged with globalAttributes
	ObsoleteAttributes     []string
	ParentContentTypes     HtmlElementType
	ParentTags             []string
	ExcludeParentTags      []string
}

func (hei *HtmlElementInfo) GetAttributeStatus(attrName string) AttrStatus {
	if attrName == "" {
		return ASUnknown
	}

	attrNameLower := strings.ToLower(attrName)

	if sorted_contains(hei.ObsoleteAttributes, attrNameLower) {
		return ASDeprecated
	}

	if sorted_contains(hei.Attributes, attrNameLower) {
		return ASValid
	}

	return ASUnknown
}

func (hei *HtmlElementInfo) IsValidParent(parentTagName string) bool {
	if parentTagName == "" {
		return true // no parent is always valid here
	}

	parentTagNameLower := strings.ToLower(parentTagName)

	// Check if the parent is in the not-allowed list
	if sorted_contains(hei.ExcludeParentTags, parentTagNameLower) {
		return false
	}

	// Check if the parent is in the white list
	if sorted_contains(hei.ParentTags, parentTagNameLower) {
		return true
	}

	// Finally, check if the content type is allowed
	if hei.ParentContentTypes == HETNone {
		return false
	}

	parentInfo := GetElementInfo(parentTagNameLower)
	if parentInfo == nil {
		if strings.Contains(parentTagName, ":") {
			return true // assume it's a custom defined element
		}

		return false
	}

	if (hei.ParentContentTypes & parentInfo.PermittedChildrenTypes) != 0 {
		return true
	}

	return false
}

func (hei *HtmlElementInfo) setPermittedChildrenTags(tags string) {
	hei.PermittedChildrenTags = convertSemicolonDelimited(tags)
}

func (hei *HtmlElementInfo) setObsoleteAttributes(attrs string) {
	hei.ObsoleteAttributes = convertSemicolonDelimited(attrs)
}

func (hei *HtmlElementInfo) setParentTags(tags string) {
	hei.ParentTags = convertSemicolonDelimited(tags)
}

func (hei *HtmlElementInfo) setExcludeParentTags(tags string) {
	hei.ParentTags = convertSemicolonDelimited(tags)
}

func (hei *HtmlElementInfo) setAttributes(attrs []string) {
	if len(attrs) == 0 {
		hei.Attributes = globalAttributes
	} else {
		hei.Attributes = union(attrs, globalAttributes)

	}
}

// GetElementInfo returns the HtmlElementInfo for this tag
func GetElementInfo(tagName string) *HtmlElementInfo {
	if tagName == "" {
		return nil
	}

	elem, exist := elemsInfo[tagName]
	if exist {
		return &elem
	}
	return nil
}

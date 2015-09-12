package htmlparser

import (
	"bytes"
	"html"
	"strings"
	//"fmt"
	"testing"
)

func Test__SimpleSegments(t *testing.T) {
	testSegments(t, true, []string{
		"",
		"text only",
		"text only with &gt; entities",
		"<b></b>",
		"<b>bold</b>",
		"a<b>bold</b>",
		"a<b>bold</b>b",
		"<b>bold<i>italic-bold</i>bold</b>",
	})
}

func Test_OptionalClosing(t *testing.T) {
	testSegments(t, true,
		[]string{
			"<p>",
			"a<p>",
			"a<p>b",
			"a<p>b</p>c",
			"a<p>b<br>c</p>",
			"a<p>b<br>c",
			"a<p>b</p><p>c</p>",
			"a<p>b</p><p>c",
			"a<p>b<p>c",
			"a<p>b<b>c</b>d</p><p>e</p>",
			"a<p>b<b>c</b>d</p><p>e",
			"a<p>b<b>c</b>d<p>e",
		})
}

func Test_OptionalClosingWithBlockElement(t *testing.T) {
	testSegments(t, true, []string{
		"<ul><li><div>a</div></li></ul>",
		"<ul><li>a<div>b</div></li></ul>",
		"<ul><li><div>a</div>b</li></ul>",
		"<ul><li><div>a</div></ul>",
		"<ul><li>a<div>b</div></ul>",
		"<ul><li><div>a</div>b</ul>",
	})
}

func Test_Scripts(t *testing.T) {
	testSegments(t, true, []string{
		"<script src=abc></script>",
		"<script>if(a<b) c;</script>",
		"abc<script>if(a<b) c;</script>def",
	})
}

func Test_Comments(t *testing.T) {
	testSegments(t, true, []string{
		"<!-- comment -->",
		"<!-- this is <b>a long comment</b> -->",
		"abcd<!-- comments <b>bold</bold> -->ghij",
	})
}

func Test_Table(t *testing.T) {
	testSegments(t, true, []string{
		"<table> <tr> <td>a</td> </table>",
		"<table> <tr> <td>a</table>",
		"<table> <tr> <td><p>a</table>",
	})
}

func Test_CompleteHtml(t *testing.T) {
	testSegments(t, true, []string{
		"<html><head><title>title</title></head><body>body</body></html>",
		"<html><head><title>hello</title><body>body",
		`<!doctype html public "-//w3c//dtd html 4.0 strict//en"><html><head><title>title</title></head><body>body</body></html>`,
	})
}

func Test_SingleTags(t *testing.T) {
	testSegments(t, true, []string{
		"<br>",
		"<hr>",
		"<br/>",
		"<br />",
		"< br />",
		"< br / >",
		"<hr size=1>",
		"<hr size=\"1\">",
		"<hr size=1/>",
		"<hr size=\"1\"/>",
	})
}

func Test_Attributes(t *testing.T) {
	testSegments(t, true, []string{
		"<a href></a>",
		"<font size=1 face=verdana>a</font>",
		"<font size=\"1\" face=verdana>a</font>",
		"<font size=1 face=\"verdana\">a</font>",
		"<font size=\"1\" face=\"verdana\">a</font>",
	})

}

func Test_StyleTag(t *testing.T) {
	testSegments(t, true, []string{
		"<head><style>.a { background-url: 'ab<>c.jpg'; }</style></head>",
	})
}

func Test_InvalidSegments(t *testing.T) {
	testSegments(t, false, []string{
		"<",
		"<b",
		"<hr",
		"<>",
		"< >",
		" < > ",
		"<b>",
		"a<b>b<i>c</b>d</i>e",
		"<a href=></a>",
		"<a href=\"abc></a>",
		"<!-- missing closing",
	})
}

func Test_InnerText(t *testing.T) {
	segments := []struct {
		Item1 string
		Item2 string
	}{
		{"", "<b></b>"},
		{"a", "<b>a</b>"},
		{"a", "a<b></b>"},
		{"a", "<b></b>a"},
		{"abc", "a<b>b</b>c"},
		{"ac", "a<!--b-->c"},
	}

	for _, segment := range segments {

		var parser = NewParser(segment.Item2)
		if !parser.Parse(nil, nil, nil) {
			t.Error()
		}
		if segment.Item1 != parser.InnerText {
			t.Error(segment.Item1)
		}

	}
}

func Test_PreserveComments(t *testing.T) {
	segment := "a<b>b<!--comment-->c</b>d"
	parser := NewParser(segment)
	parser.SkipComments = false
	if !parser.Parse(nil, nil, nil) {
		t.Error()
	}
	if parser.InnerText != "ab<!--comment-->cd" {
		t.Error()
	}

}

func Test_ComplexHtml(t *testing.T) {
	parser := NewParser(googleHomepage)
	parser.Parse(nil, nil, nil)
	if !parser.HasValidSyntax {
		t.Error()
	}
}

func Test_CustomInnerText(t *testing.T) {
	segment := "a<b>b</b>c<!--comment-->d<p>e"

	n := bytes.NewBufferString("")

	parser := NewParser(segment)

	parser.Parse(func(text string, he *HtmlElement) {
		n.WriteString(text)
	}, nil, nil)

	if n.String() != "abcde" {
		t.Error()
	}

}

func Test_UrlAttribute(t *testing.T) {
	segment := `<b> <link rel="alternate" type="application/rss+xml" title="M-Shaped Brain &raquo; Feed" href="http://blog.calbucci.com/feed/" /> </b>`

	foundIt := false
	parser := NewParser(segment)

	parser.Parse(nil, func(e *HtmlElement, isEmpty bool) {
		t.Logf("E: %v (attr=%v)\n", e.TagName, e.Attributes)
		if e.TagName == "link" {
			if len(e.Attributes) != 4 {
				t.Error()
			}
			if title, _ := e.GetAttributeValue("title"); title != "M-Shaped Brain » Feed" {
				t.Error()
			}
			if href, _ := e.GetAttributeValue("href"); href != "http://blog.calbucci.com/feed/" {
				t.Error()
			}

			foundIt = true
		}
	}, nil)

	if !foundIt {
		t.Error()
	}

}

func Test_FindRSSFeed(t *testing.T) {
	rssFeed := ""
	parser := NewParser(blogPost)

	parser.Parse(nil, func(e *HtmlElement, isEmpty bool) {
		if e.TagName == "link" {

			if ty, _ := e.GetAttributeValue("type"); ty == "application/rss+xml" {
				t.Logf("rss-e: %v %v\n", e.TagName, e.Attributes)
				rssFeed, _ = e.GetAttributeValue("href")
				parser.Stop()
			}
		}
	}, nil)

	t.Logf("rssFeed=%v\n", rssFeed)
	if rssFeed != "http://blog.calbucci.com/feed/" {
		t.Error()
	}

}

func Test_Idempotent(t *testing.T) {
	baseHtml := blogPost
	html1 := parseAndSerialize(baseHtml)
	html2 := parseAndSerialize(html1)
	html3 := parseAndSerialize(html2)

	if html1 != html2 {

		max := len(html1)
		if max > len(html2) {
			max = len(html2)
		}
		for i := 0; i < max; i++ {
			if html1[i] != html2[i] {
				i -= 20
				if i < 0 {
					i = 0
				}
				e := i + 30
				if e > max {
					e = max
				}
				t.Logf("Mismatch1: %v\n", html1[i:e])
				t.Logf("Mismatch2: %v\n", html2[i:e])
				break
			}
		}

		t.Error()
	}
	if html2 != html3 {
		t.Error()
	}
}

func parseAndSerialize(origHtml string) string {
	parser := NewParser(origHtml)

	parser.PreserveCRLFTab = false

	n := bytes.NewBufferString("")

	parser.Parse(func(text string, parent *HtmlElement) {
		escaped := html.EscapeString(text)
		n.WriteString(escaped)
	}, func(parent *HtmlElement, isEmptyTag bool) {
		n.WriteString(parent.GetOpenTag(false, false))
	}, func(closeTag string) {
		n.WriteString("</" + closeTag + ">")
	})

	return n.String()
}

func Test_FindOpenGraphTags(t *testing.T) {
	parser := NewParser(blogPost)

	tags := make(map[string]string)

	parser.Parse(nil, func(element *HtmlElement, isEmptyTag bool) {
		if element.TagName == "meta" {
			ogName, _ := element.GetAttributeValue("property")
			if ogName == "" || !strings.HasPrefix(ogName, "og:") {
				return
			}
			ogValue, _ := element.GetAttributeValue("content")
			tags[ogName] = ogValue
		}
	}, nil)

	if !parser.HasValidSyntax {
		t.Error()
	}

	if v, _ := tags["og:type"]; v != "article" {
		t.Error()
	}

	if v, _ := tags["og:url"]; v != "http://blog.calbucci.com/2015/01/27/attention-cannibalization/" {
		t.Error()
	}

}

func testSegments(t *testing.T, result bool, segments []string) {
	for _, segment := range segments {
		t.Logf("Processing: %v\n", segment)
		parser := NewParser(segment)
		if parser.Parse(nil, nil, nil) != result {
			t.Errorf("Failed to parse segment: " + segment)
		}
	}
}

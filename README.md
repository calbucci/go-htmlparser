# go-htmlparser
Events-based HTML 5.0 compliant parser in Go (SAX-style parsing)

## Typical Scenarios
- Use it to scrape pieces of HTML
- Detect META / LINK tags (e.g. Open Graph tags)
- Optimize the output HTML (remove whitespace, clear empty tags)
- Detect HTML syntax errors and notify developers
- Extract text from the HTML


## Sample

### Get the RSS Feed of a website

```go
	rssFeed := ""
	parser := NewParser(htmlContent)

	parser.Parse(nil, func(e *HtmlElement, isEmpty bool) {
		if e.TagName == "link" {

			if ty,_ := e.GetAttributeValue("type"); ty == "application/rss+xml" {
				t.Logf("rss-e: %v %v\n", e.TagName, e.Attributes)
				rssFeed,_ = e.GetAttributeValue("href")
				parser.Stop()
			}
		}
	}, nil)
	
	fmt.Println(rssFeed)
```

### Remove whitespaces

```go
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

	newHtml := n.String()
```



## Questions



## Contributors

- HtmlParser was originally created by *Marcelo Calbucci* ([blog.calbucci.com](http://blog.calbucci.com) | [@calbucci](http://twitter.com/calbucci))


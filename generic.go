package htmlparser

import (
	"bytes"
	"sort"
	"strings"
	"unicode"
)

// union add all the elements from slice2 to slice1 if not present
func union(slice1, slice2 []string) []string {
	if slice1 == nil {
		slice1 = make([]string, len(slice2))
		copy(slice1, slice2)
		sort.Strings(slice1)
		return slice1
	}
	for _, e2 := range slice2 {
		found := false
		for _, e1 := range slice1 {
			if e1 == e2 {
				found = true
				break
			}
		}
		if !found {
			slice1 = append(slice1, e2)
		}
	}
	sort.Strings(slice1)
	return slice1
}

func sorted_contains(slice []string, element string) bool {
	if slice == nil || len(slice) == 0 {
		return false
	}
	return sort.SearchStrings(slice, element) >= 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func convertSemicolonDelimited(text string) []string {
	if len(text) > 0 {
		strList := strings.Split(text, ";")
		if len(strList) > 0 {
			for i, s := range strList {
				strList[i] = strings.ToLower(s)
			}
			sort.Strings(strList)
			return strList
		}
	}
	return nil
}

func runesLastIndex(runes []rune, r rune) int {

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == r {
			return i
		}
	}
	return -1
}

func runesIndex(runes []rune, r rune) int {
	for i, v := range runes {
		if v == r {
			return i
		}
	}
	return -1
}

func runesIndexRunesStart(runes []rune, sub []rune, start int) int {

	max := len(runes) - len(sub)

	for ; start < max; start++ {

		match := true
		for i := 0; i < len(sub); i++ {
			if runes[start+i] != sub[i] {
				match = false
				break
			}
		}
		if match {
			return start
		}

	}
	return -1
}

func trimInBetween(str string) string {
	if str == "" {
		return str
	}

	n := bytes.NewBufferString("")

	lastSpace := true

	for _, r := range str {

		if unicode.IsSpace(r) || unicode.IsControl(r) {
			if lastSpace {
				continue
			}
			lastSpace = true
			n.WriteRune(' ')
			continue
		}
		n.WriteRune(r)
	}
	return n.String()
}

func hasContent(text string) bool {
	if len(text) == 0 {
		return false
	}

	for i, r := range text {
		if !unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

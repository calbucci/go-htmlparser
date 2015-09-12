package htmlparser

import (
	"strconv"
	"testing"
)

func Test_Union(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"a", "c"}

	c := union(a, b)

	if len(c) != 3 {
		t.Error()
	}
}

func Test_Sorted_contains(t *testing.T) {
	a := []string{"a", "b", "c"}

	for _, v := range a {
		if !sorted_contains(a, v) {
			t.Error(v)
		}
	}

	if sorted_contains(a, "d") {
		t.Error("d")
	}
}

func Test_Contains(t *testing.T) {
	a := []string{"a", "b", "c"}

	for _, v := range a {
		if !sorted_contains(a, v) {
			t.Error(v)
		}
	}

	if sorted_contains(a, "d") {
		t.Error("d")
	}
}

func Test_convertSemicolonDelimited(t *testing.T) {
	r := convertSemicolonDelimited("")
	if r != nil {
		t.Error()
	}

	r = convertSemicolonDelimited("a")
	if len(r) != 1 || r[0] != "a" {
		t.Error()
	}

	r = convertSemicolonDelimited("a;a")
	if len(r) != 2 || r[0] != "a" {
		t.Error()
	}

	r = convertSemicolonDelimited("b;a")
	if len(r) != 2 || r[0] != "a" || r[1] != "b" {
		t.Error()
	}

}

func Test_runesLastIndex(t *testing.T) {
	if runesLastIndex([]rune(""), 'a') != -1 {
		t.Error()
	}

	if runesLastIndex([]rune("a"), 'a') != 0 {
		t.Error()
	}

	if runesLastIndex([]rune("bac"), 'a') != 1 {
		t.Error()
	}

	if runesLastIndex([]rune("bba"), 'a') != 2 {
		t.Error()
	}

	if runesLastIndex([]rune("abab"), 'a') != 2 {
		t.Error()
	}

	if runesLastIndex([]rune("defg"), 'a') != -1 {
		t.Error()
	}
}

func Test_runesIndex(t *testing.T) {
	if runesIndex([]rune(""), 'a') != -1 {
		t.Error()
	}

	if runesIndex([]rune("a"), 'a') != 0 {
		t.Error()
	}

	if runesIndex([]rune("bac"), 'a') != 1 {
		t.Error()
	}

	if runesIndex([]rune("bba"), 'a') != 2 {
		t.Error()
	}

	if runesIndex([]rune("abab"), 'a') != 0 {
		t.Error()
	}

	if runesIndex([]rune("defg"), 'a') != -1 {
		t.Error()
	}
}

func Test_runesIndexRunesStart(t *testing.T) {
	if runesIndexRunesStart([]rune(""), []rune(""), 0) != -1 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abc"), []rune(""), 0) != -1 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abc"), []rune("d"), 0) != -1 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abc"), []rune("def"), 0) != -1 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abc"), []rune("abd"), 0) != -1 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abc"), []rune("a"), 0) != 0 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abc"), []rune("c"), 0) != 2 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abc"), []rune("abc"), 0) != 0 {
		t.Error()
	}

	if runesIndexRunesStart([]rune("abab"), []rune("ab"), 0) != 0 {
		t.Error()
	}

	if r := runesIndexRunesStart([]rune("abab"), []rune("ab"), 1); r != 2 {
		t.Error(strconv.Itoa(r))
	}

}

func Test_trimInBetween(t *testing.T) {

	if trimInBetween("") != "" {
		t.Error()
	}

	if trimInBetween("abc") != "abc" {
		t.Error()
	}

	if trimInBetween(" abc ") != " abc " {
		t.Error()
	}

	if trimInBetween("a b c") != "a b c" {
		t.Error()
	}

	if trimInBetween("a  b  c") != "a b c" {
		t.Error()
	}

	if trimInBetween("a\nb\nc") != "a b c" {
		t.Error()
	}

	if r := trimInBetween("a\n\nb \n c"); r != "a b c" {
		t.Error(r)
	}
}

func Test_hasContent(t *testing.T) {
	if hasContent("") {
		t.Error()
	}

	if hasContent(" ") {
		t.Error()
	}

	if hasContent("\r") {
		t.Error()
	}

	if hasContent("\r\n\t ") {
		t.Error()
	}

	if !hasContent("a") {
		t.Error()
	}

	if !hasContent(" a") {
		t.Error()
	}

	if !hasContent("a ") {
		t.Error()
	}

	if !hasContent("\t \n a") {
		t.Error()
	}

}

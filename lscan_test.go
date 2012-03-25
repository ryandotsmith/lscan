package lscan

import (
	"strings"
	"testing"
)

var str = `hello=world name="ryan\"smith" distance=1.123 desc="hi=there" time="2012-03-21 10:18:20 -0700"`

func TestParseSimple(t *testing.T) {
	in := strings.NewReader(str)
	m := Parse(in)

	actual := m["hello"]
	expected := "world"
	if actual != expected {
		t.Errorf("\n e(%v) \n a(%v)", expected, actual)
	}
}

func TestParseQuoted(t *testing.T) {
	in := strings.NewReader(str)
	m := Parse(in)

	actual := m["name"]
	expected := `"ryan\"smith"`
	if actual != expected {
		t.Errorf("\n e(%v) \n a(%v)", expected, actual)
	}
}

func TestParseNum(t *testing.T) {
	in := strings.NewReader(str)
	m := Parse(in)

	actual := m["distance"]
	expected := `1.123`
	if actual != expected {
		t.Errorf("\n e(%v) \n a(%v)", expected, actual)
	}
}

func TestParseTime(t *testing.T) {
	in := strings.NewReader(str)
	m := Parse(in)

	actual := m["time"]
	expected := `"2012-03-21 10:18:20 -0700"`
	if actual != expected {
		t.Errorf("\n e(%v) \n a(%v)", expected, actual)
	}
}

func TestParseEqchr(t *testing.T) {
	in := strings.NewReader(str)
	m := Parse(in)

	actual := m["desc"]
	expected := `"hi=there"`
	if actual != expected {
		t.Errorf("\n e(%v) \n a(%v)", expected, actual)
	}
}

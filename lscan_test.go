package lscan

import (
	"strings"
	"testing"
)

var str = `hello=world name="ryan\"smith" distance=1.123 desc="hi=there" time="2012-03-21 10:18:20 -0700"`
var aStr = ` [590-1]  [PINK] LOG:  checkpoint complete: wrote 0 buffers (0.0%); 0 transaction log file(s) added, 0 removed, 1 recycled; write=0.001 s, sync=0.000 s, total=0.008 s; sync files=0, longest=0.000 s, average=0.000 s`

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

func TestParseAStr(t *testing.T) {
	in := strings.NewReader(aStr)
	m := Parse(in)

	actual := m["write"]
	expected := `0.001`
	if actual != expected {
		t.Errorf("\n e(%v) \n a(%v)", expected, actual)
	}

	if len(m) != 6 {
		t.Errorf("\n expected 6 pairs \n received %v", len(m))
	}
}

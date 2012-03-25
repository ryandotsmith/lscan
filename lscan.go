package lscan

import (
	"io"
	"text/scanner"
)

var s scanner.Scanner

func Parse(r io.Reader) (kvmap map[string]interface{}) {
	s.Init(r)
	s.Whitespace = 0 // don't skip whitespace
	tok := s.Scan()
	kvmap = make(map[string]interface{})

	key := ""
	val := ""
	ink := true // we always start with the key
	inv := false

	for tok != scanner.EOF {
		tex := s.TokenText()
		if tex == `=` {
			ink = false
			inv = true
			tok = s.Scan()
			continue
		}
		if tex == ` ` {
			writeTup(kvmap, key, val)
			key = ""
			val = ""
			ink = true
			inv = false
			tok = s.Scan()
			continue
		}
		if ink {
			key += tex
		} else if inv {
			val += tex
		}

		tok = s.Scan()
		// We are done. :( Let us write whatever we have left
		// to the map
		if tok == scanner.EOF {
			writeTup(kvmap, key, val)
		}
	}
	return
}

func writeTup(m map[string]interface{}, k string, v string) {
	if len(k) > 0 && len(v) > 0 {
		m[k] = v
	}
}

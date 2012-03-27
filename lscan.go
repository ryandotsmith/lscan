package lscan

import (
	"io"
	"text/scanner"
)

var s scanner.Scanner

func Parse(r io.Reader) map[string]string {
	s.Init(r)
	s.Whitespace = 0 // don't skip whitespace
	kvmap := make(map[string]string)

	key := ""
	val := ""
	ink := true // if we are not in key we are in val

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tex := s.TokenText(); tex {
		case `=`:
			ink = false
		case ` `:
			writeTup(kvmap, key, val)
			key = ""
			val = ""
			ink = true
		default:
			if ink {
				key += tex
			} else {
				val += tex
			}
		}
	}
	writeTup(kvmap, key, val)
	return kvmap
}

func writeTup(m map[string]string, k string, v string) {
	if len(k) > 0 && len(v) > 0 {
		m[k] = v
	}
}

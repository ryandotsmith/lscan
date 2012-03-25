package lscan

import (
	"fmt"
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
		fmt.Printf("handle=%v ", tex)
		if tex == `=` {
			fmt.Printf("split ")
			ink = false
			inv = true
			tok = s.Scan()
			continue
		}
		if tex == ` ` {
			fmt.Printf("next ")
			kvmap[key] = val
			key = ""
			val = ""
			ink = true
			inv = false
			tok = s.Scan()
			continue
		}
		if ink {
			key += tex
			fmt.Printf("key=%v ", key)
		} else if inv {
			val += tex
			fmt.Printf("val=%v ", val)
		}

		tok = s.Scan()
		// We are done. :( Let us write whatever we have left
		// to the map
		if tok == scanner.EOF {
			kvmap[key] = val
		}
	}
	return
}

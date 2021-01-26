//The standard library’s strings package provides many useful string-related functions. Here are some examples to give you a sense of the package.
package main

import s "strings"
import "fmt"

//We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()
	//Not part of strings, but worth mentioning here, are the mechanisms for getting the length of a string in bytes and getting a byte by index.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}

//Note that len and indexing above work at the byte level. Go uses UTF-8 encoded strings, so this is often useful as-is. If you’re working with potentially multi-byte characters you’ll want to use encoding-aware operations. See strings, bytes, runes and characters in Go for more information.

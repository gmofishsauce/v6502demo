package main

import "encoding/xml"
import "fmt"
import "log"
import "os"
import "strings"

func main() {
	in := os.Stdin
	name := "standard input"
	if len(os.Args) > 1 {
		name = os.Args[1]
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("opening %s: %s", name, err)
		}
		defer f.Close()
		in = f
	}

	d := xml.NewDecoder(in)

	// Check for the HTML strict directive and then
	// enter recursive descent for the document.
	t := get(d)
	bytes, ok := t.(xml.Directive)
	if !ok {
		fail("XML Strict directive", t)
	}
	content := string(bytes)
	if !strings.Contains(content, "XHTML 1.0 Strict") {
		log.Fatalf("not an XHTML 1.0 Strict document")
	}
	
	html(d)
}

// HTML start and end tags surround head and body
func html(d *xml.Decoder) {
	getStart(d, "html")
	head(d)
	body(d)
	getEnd(d, "html")
}

func head(d *xml.Decoder) {
	getStart(d, "head")
	title(d)
	getEnd(d, "head")
}

func body(d *xml.Decoder) {
	getStart(d, "body")
	getEnd(d, "body")
}

func title(d *xml.Decoder) {
	getStart(d, "title")
	t := get(d)
	cd, ok := t.(xml.CharData)
	if !ok {
		fail("title text", t)
	}
	fmt.Printf("# %s\n", string(cd))
	getEnd(d, "title")
}

// === End of recursive descent parser ===
// Utility functions below

// Get the next token, skipping whitespace
func get(d *xml.Decoder) xml.Token {
	var result xml.Token
	var err error
	for {
		result, err = d.Token()
		if err != nil {
			log.Fatal(err)
		}
		if !isWhiteSpace(result) {
			break
		}
	}
	return result
}

func getStart(d *xml.Decoder, expected string) {
	t := get(d)
	start, ok := t.(xml.StartElement)
	if !ok || start.Name.Local != expected {
		fail(expected, t)
	}
}

func getEnd(d *xml.Decoder, expected string) {
	t := get(d)
	end, ok := t.(xml.EndElement)
	if !ok || end.Name.Local != expected {
		fail(expected, t)
	}
}

func isWhiteSpace(t xml.Token) bool {
	cd, ok := t.(xml.CharData)
	if !ok {
		// not char data => not whitespace
		return false
	}

	s := string(cd)
	if len(strings.TrimSpace(s)) != 0 {
		return false
	}
	return true
}

func fail(expected string, got any) {
	log.Fatalf("%s expected, got %v", expected, got)
}


package main

import "encoding/xml"
import "fmt"
import "log"
import "os"
import "reflect"
import "strings"
import "time"

var debug bool = false

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
	
	result := parse(d, "html")
	walk(result, 0)
	emitString("\nWritten by mkmd %s\n", time.Now().String())
}

// Node structures produced by the recursive descent parser
type node struct {
	name string
	text string
	attrs []xml.Attr
	children []*node
	processed bool   // development/debugging hack only
}

// The recursive descent parser. Just builds a tree of nodes.
func parse(d *xml.Decoder, expected string) *node {
	result := getStart(d, expected)
	loop:
	for {
		t := get(d)
		switch inst := t.(type) {
		case xml.StartElement:
			s := inst.Name.Local
			unget(t)
			result.children = append(result.children, parse(d, s))
		case xml.EndElement:
			unget(t)
			break loop;
		case xml.CharData:
			result.children = append(result.children, &node{name: "CharData", text: string(inst)})
		default:
			// Directive or something: "type not expected"
			// Improve the parser if this ever occurs
			fail("type not", inst)
		}
	}
	getEnd(d, expected)
	return result
}

// The second pass. Walks the tree calling start and end emit for each node
func walk(np *node, indent int) {
	startEmit(np)	
	for _, c := range(np.children) {
		walk(c, 1 + indent)
	}
	endEmit(np)
	if !np.processed {
		fmt.Printf("\nUNSUPPORTED: %s\n", np.name)
	}
}

// Markdown emitters

func startEmit(np *node) {
	switch np.name {
	case "CharData":
		// emitString(np.text)		
		emitString("text")
		np.processed = true
	case "a":
		emitString("[")
	case "body":
		np.processed = true
	case "div":
		// For now we don't do anything with divs.
		np.processed = true
	case "h1", "h2", "h3", "h4", "h5", "h6":
		nHashes := np.name[1] - '0'
		emitString(strings.Repeat("#", int(1 + nHashes)))
		np.processed = true
	case "head":
		np.processed = true
	case "html":
		np.processed = true
	case "img":
		emitString("![an image](%s)", getAttrValue(np, "src"))
		//emitString("image here")
		np.processed = true
	case "p":
		emitString("\n")
		np.processed = true
	case "span":
		np.processed = true
	case "td":
		emitString("|")
		np.processed = true
	case "title":
		emitString("\n# ")
		np.processed = true
	case "tr":
		emitString("\n\n")
	}
}

func endEmit(np *node) {
	switch np.name {
	case "a":
		emitString("](%s)", "href here") // getAttrValue(np, "href"))
		np.processed = true
	case "table":
		emitString(strings.Repeat("|:---:", len(np.children[0].children)))
		emitString("|\n")
		np.processed = true
	case "tr":
		emitString("|\n")
		np.processed = true
	}
}

func emitString(format string, args ...any) {
	fmt.Printf(format, args...)
}

func getAttrValue(np *node, toFind string) string {
	for _, attr := range(np.attrs) {
		if attr.Name.Local == toFind {
			return attr.Value
		}
	}
	fail(fmt.Sprintf("attribute %s for node %s", toFind, np.name), "nothing")
	return "internal error"
}

// Token getters

var pushback xml.Token = nil

// Get the next token, skipping whitespace
// and comment tokens.
func get(d *xml.Decoder) xml.Token {
	var result xml.Token
	if pushback != nil {
		result = pushback
		pushback = nil
		debugToken("get", result)
		return result
	}

	var err error
	for {
		result, err = d.Token()
		if err != nil {
			log.Fatalf("while reading: %s", err)
		}
		// Is it white space or comment or etc?
		if !skip(result) {
			break
		}
	}

	// Tokens returned by d.Token()
	// are ephemeral so we need CopyToken()
	// to allow pushback
	debugToken("get", result)
	return xml.CopyToken(result)
}

func unget(t xml.Token) {
	if pushback != nil {
		log.Fatalf("too many pushbacks")
	}
	debugToken("unget", t)
	pushback = t
}

// Get (expect) a specific start tag
func getStart(d *xml.Decoder, expected string) *node {
	t := get(d)
	start, ok := t.(xml.StartElement)
	if !ok || start.Name.Local != expected {
		fail(expected, t)
	}
	return &node{name: start.Name.Local, attrs: start.Attr}
}

func getEnd(d *xml.Decoder, expected string) {
	t := get(d)
	end, ok := t.(xml.EndElement)
	if !ok || end.Name.Local != expected {
		fail(expected, t)
	}
}

// Utilities

// Return true if the token is ignoreable
// (whitespace or HTML comment)
func skip(t xml.Token) bool {
	_, ok := t.(xml.Comment)
	if ok { // comment: skip
		return true
	}
	cd, ok := t.(xml.CharData)
	if ok {
		s := string(cd)
		if len(strings.TrimSpace(s)) == 0 {
			// whitespace: skip
			return true
		}
	}

	return false
}

// Slightly intelligent stringify
func anyToString(data any) string {
	switch inst := data.(type) {
	case xml.StartElement:
		return inst.Name.Local
	case xml.EndElement:
		return inst.Name.Local
	}

    // Use reflection to check if data is really a slice of bytes ([]byte)
    valueType := reflect.TypeOf(data)
    if valueType.Kind() != reflect.Slice || valueType.Elem().Kind() != reflect.Uint8 {
        return fmt.Sprintf("%v", data)
    }

    // Use reflection to convert the []byte to a string
    valueBytes := reflect.ValueOf(data)
    return string(valueBytes.Bytes())
}

func debugToken(msg string, t xml.Token) {
	if debug {
		fmt.Printf("  [%s [%T] %s]\n", msg, t, anyToString(t))
	}
}

func fail(expected string, got any) {
	log.Fatalf("%s expected, got (%T) %s", expected, got, anyToString(got))
}


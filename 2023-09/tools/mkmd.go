package main

import "encoding/xml"
import "fmt"
import "log"
import "os"
import "reflect"
import "strings"
import "time"

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
	emit("\nWritten by mkmd %s\n", time.Now().String())
	getEnd(d, "html")
}

// I wrote the code that cleaned up the HTML files we're reading,
// so I know the only thing that can be in the head section is
// a single title tag.
func head(d *xml.Decoder) {
	getStart(d, "head")
	title(d)
	getEnd(d, "head")
}

func body(d *xml.Decoder) {
	bodyTag(d, "body")
}

func bodyTag(d *xml.Decoder, expected string) {
	getStart(d, expected)
	loop:
	for {
		t := get(d)
		switch inst := t.(type) {
		case xml.StartElement:
			s := inst.Name.Local
			unget(t)
			switch s {
			case "a":
				a(d)
			case "br":
				br(d)
			case "div":
				div(d)
			case "form":
				form(d)
			case "h1":
				h1(d)
			case "h2":
				h2(d)
			case "h3":
				h3(d)
			case "h4":
				h4(d)
			case "img":
				img(d)
			case "li":
				li(d)
			case "p":
				p(d)
			case "pre":
				pre(d)
			case "script":
				script(d)
			case "span":
				span(d)
			case "sup":
				sup(d)
			case "table":
				table(d)
			case "td":
				td(d)
			case "tr":
				tr(d)
			case "ul":
				ul(d)
			default:
				fail("tag not", t)
			}
		case xml.EndElement:
			unget(t)
			break loop;
		case xml.CharData:
			emit(string(inst))
		default:
			// Directive or something
			// "type not expected"
			// Improve the parser
			fail("type not", inst)
		}
	}
	getEnd(d, expected)
}

func a(d *xml.Decoder) {
	bodyTag(d, "a")
}

func br(d *xml.Decoder) {
	bodyTag(d, "br")
}

func div(d *xml.Decoder) {
	bodyTag(d, "div")
}

func form(d *xml.Decoder) {
	bodyTag(d, "form")
}

func h1(d *xml.Decoder) {
	bodyTag(d, "h1")
}

func h2(d *xml.Decoder) {
	bodyTag(d, "h2")
}

func h3(d *xml.Decoder) {
	bodyTag(d, "h3")
}

func h4(d *xml.Decoder) {
	bodyTag(d, "h4")
}

func img(d *xml.Decoder) {
	bodyTag(d, "img")
}

func li(d *xml.Decoder) {
	bodyTag(d, "li")
}

func p(d *xml.Decoder) {
	bodyTag(d, "p")
}

func pre(d *xml.Decoder) {
	bodyTag(d, "pre")
}

func script(d *xml.Decoder) {
	bodyTag(d, "script")
}

func span(d *xml.Decoder) {
	bodyTag(d, "span")
}

func sup(d *xml.Decoder) {
	bodyTag(d, "sup")
}

func table(d *xml.Decoder) {
	bodyTag(d, "table")
}

func td(d *xml.Decoder) {
	bodyTag(d, "td")
}

func tr(d *xml.Decoder) {
	bodyTag(d, "tr")
}

func ul(d *xml.Decoder) {
	bodyTag(d, "ul")
}

func title(d *xml.Decoder) {
	getStart(d, "title")
	t := get(d)
	cd, ok := t.(xml.CharData)
	if !ok {
		fail("title text", t)
	}
	emit("# %s\n", string(cd))
	getEnd(d, "title")
}

// Utility functions below

func emit(format string, args ...any) {
	fmt.Printf(format, args...)
}

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
	fmt.Printf("  [%s [%T] %s]\n", msg, t, anyToString(t))
}

func fail(expected string, got any) {
	log.Fatalf("%s expected, got (%T) %s", expected, got, anyToString(got))
}


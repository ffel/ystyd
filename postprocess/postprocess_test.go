package postprocess

import (
	"fmt"
	"testing"

	"log"
	"launchpad.net/goyaml"
)

var yaml string = `site:
- in: index.md
  menu: Home
  out: index.html
- in: over.md
  menu: About
  out: about.html
- in: contact.md
  menu: Contact
  out: contact.html
nav:
  menu: <nav>{{.Menu}}</nav>
  active: <li class="active"><a href="{{.Href}}">{{.Label}}</a></li>
  inactive: <li><a href="{{.Href}}">{{.Label}}</a></li>
`

func ExamplePostProcess() {
	d := NewSite()
	err := d.Read(yaml)

	if err != nil {
		log.Fatal(err)
	}

	page, err := d.PostProcess("index.html", "<site>\n{{.Nav}}\n</site>")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(page)

	// Output:
	// <site>
	// <nav><li class="active"><a href="index.html">Home</a></li>
	// <li><a href="about.html">About</a></li>
	// <li><a href="contact.html">Contact</a></li>
	// </nav>
	// </site>
}

func TestUnmarshal(t *testing.T) {

	d := Site{}
	err := goyaml.Unmarshal([]byte(yaml), &d)
	if err != nil {
		panic(err)
	}

	read := fmt.Sprintf("%v", d.Pages)
	exp := "[{index.md Home index.html} {over.md About about.html} {contact.md Contact contact.html}]"
	if read != exp {
		t.Errorf("error:, %q != %q", read, exp)
	}

	read = fmt.Sprintf("%v", d.Menu)
	exp = "{<nav>{{.Menu}}</nav> <li class=\"active\"><a href=\"{{.Href}}\">{{.Label}}</a></li> <li><a href=\"{{.Href}}\">{{.Label}}</a></li>}"
	if read != exp {
		t.Errorf("error:, %q != %q", read, exp)
	}
}

func TestRead(t *testing.T) {
	d := NewSite()

	err := d.Read(yaml)

	if err != nil {
		t.Errorf("error: non nill error %v", err)
	}
}

func TestCreate(t *testing.T) {
	d := NewSite()
	err := d.Read(yaml)

	if err != nil {
		t.Errorf("error: non nill error %v", err)
	}

	got, err := d.create("index.html")
	exp := `<nav><li class="active"><a href="index.html">Home</a></li>
<li><a href="about.html">About</a></li>
<li><a href="contact.html">Contact</a></li>
</nav>`

	if err != nil {
		t.Errorf("error: non nill error %v", err)
	}

	// with %s instead of %q you get rid of the escaped quotes
	if got != exp {
		t.Errorf("error:\n%s\n\t!=\n%s", got, exp)
	}
}

func TestCreate_noRead(t *testing.T) {
	d := NewSite()

	// forgot to Read

	got, err := d.create("index.html")
	exp := ""

	if err != nil {
		t.Errorf("error: non nill error %v", err)
	}

	if got != exp {
		t.Errorf("error:\n%s\n\t!=\n%s", got, exp)
	}
}

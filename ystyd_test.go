package ystyd

import (
	"fmt"
	"testing"

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
  menu: <m>{{.Menu}}</m>
  active: <li class="active"><a href="{{.Href}}">{{.Label}}</a></li>
  inactive: <li><a href="{{.Href}}">{{.Label}}</a></li>
`

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
	exp = "{<m>{{.Menu}}</m> <li class=\"active\"><a href=\"{{.Href}}\">{{.Label}}</a></li> <li><a href=\"{{.Href}}\">{{.Label}}</a></li>}"
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

	read := fmt.Sprintf("%v", d.Create("index.html"))
	exp := `<m><li class="active"><a href="index.html">Home</a></li>
<li><a href="about.html">About</a></li>
<li><a href="contact.html">Contact</a></li>
</m>`

	// with %s instead of %q you get rid of the escaped quotes
	if read != exp {
		t.Errorf("error:\n%s\n\t!=\n%s", read, exp)
	}
}

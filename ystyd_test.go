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
  menu: <m>{{.menu}}</m>
  active: <li class="active"><a href="{{.href}}">{{.label}}</a></li>
  inactive: <li><a href="{{.href}}">{{.label}}</a></li>
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
	exp = "{<m>{{.menu}}</m> <li class=\"active\"><a href=\"{{.href}}\">{{.label}}</a></li> <li><a href=\"{{.href}}\">{{.label}}</a></li>}"
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
	exp := "boo"

	if read != exp {
		t.Errorf("error:, %q != %q", read, exp)
	}
}

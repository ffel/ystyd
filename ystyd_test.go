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
  menu: <ul class="nav nav-pills pull-right">{{.menu}}</ul>
  active: <li class="active"><a href="{{.href}}">{{.label}}</a></li>
  inactive: <li><a href="{{.href}}">{{.label}}</a></li>
`

func TestData(t *testing.T) {

	d := Data{}
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
	exp = "{<ul class=\"nav nav-pills pull-right\">{{.menu}}</ul> <li class=\"active\"><a href=\"{{.href}}\">{{.label}}</a></li> <li><a href=\"{{.href}}\">{{.label}}</a></li>}"
	if read != exp {
		t.Errorf("error:, %q != %q", read, exp)
	}
}

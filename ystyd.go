// Package Ystyd provides means to complement pandoc in
// generating static web sites
package ystyd

import (
	"bytes"
	"text/template"

	"launchpad.net/goyaml"
)

// struct Page contains In: input md file name, Menu: header in Site Menu
// Out: site html file name
type Page struct {
	In   string `yaml:"in,omitempty"`
	Menu string `yaml:"menu,omitempty"`
	Out  string `yaml:"out,omitempty"`
}

// struct Templates contains templates for menu wrapper, active page
// and inactive page
type Templates struct {
	Menu     string `yaml:"menu,omitempty"`
	Active   string `yaml:"active,omitempty"`
	Inactive string `yaml:"inactive,omitempty"`
}

// struct Site is the wrapper struct for yaml data
type Site struct {
	Pages []Page    `yaml:"site,omitempty"`
	Menu  Templates `yaml:"nav,omitempty"`
}

// NewSite creates an empty site structure reference
func NewSite() *Site {
	return &Site{}
}

// Read read yaml site data
func (d *Site) Read(data string) error {
	err := goyaml.Unmarshal([]byte(data), &d)
	if err != nil {
		return err
	}
	return nil
}

// Create creates the html menu for file
func (d *Site) create(file string) (string, error) {
	// template.Execute expects a writes, we use a buffer here
	var b bytes.Buffer

	// create the contents
	for _, page := range d.Pages {
		t := template.New("menuentry")

		var err error

		if file == page.Out {
			t, err = t.Parse(d.Menu.Active)
		} else {
			t, err = t.Parse(d.Menu.Inactive)

		}

		if err != nil {
			return "", err
		}

		err = t.Execute(&b, struct {
			Href  string
			Label string
		}{page.Out, page.Menu})

		if err != nil {
			return "", err
		}

		// according to the docs, err is always nil, so we ignore it
		b.WriteString("\n")
	}

	// wrap the contents in Menu
	var buff bytes.Buffer

	m := template.New("menu")
	m, err := m.Parse(d.Menu.Menu)

	if err != nil {
		return "", err
	}

	err = m.Execute(&buff, struct {
		Menu string
	}{b.String()})

	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

// PostProcess adds menu to file fname with contents f
func (d *Site) PostProcess(fname string, f string) (string, error) {

	return "", nil
}

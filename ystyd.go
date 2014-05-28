package ystyd

import "launchpad.net/goyaml"

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
	return new(Site)
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
func (d *Site) Create(file string) string {
	return d.Menu.Menu
}

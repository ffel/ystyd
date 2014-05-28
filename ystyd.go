package ystyd

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

// struct Data is the wrapper struct for yaml data
type Data struct {
	Pages []Page    `yaml:"site,omitempty"`
	Menu  Templates `yaml:"nav,omitempty"`
}

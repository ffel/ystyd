package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ffel/ystyd/postprocess"
)

func main() {
	configPtr := flag.String("config", "config.yaml", "configuration yaml file")
	siteDirPtr := flag.String("site", "./Public", "static site dir")
	postprocessPtr := flag.Bool("post", true, "postproces site creation: add site nav")

	flag.Parse()

	if *postprocessPtr {
		insertSiteNav(*siteDirPtr, *configPtr)
	}
}

// insertSiteNav inserts site navigation in {{.Nav}}
// of html files in dir according to configuration in cfile
func insertSiteNav(dir, cfile string) {
	fmt.Printf("post process %q with config file %q\n", dir, cfile)

	readConfig(cfile)
	createSite()
	addPages(dir)

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".html" {
			// read html file with {{.Nav}}
			templateBytes, err := ioutil.ReadFile(dir + "/" + f.Name())
			checkForErrors(err)

			// insert Nav menu
			html, err := site.PostProcess(f.Name(), string(templateBytes))
			checkForErrors(err)

			// write the html file
			err = ioutil.WriteFile(dir+"/"+f.Name(), []byte(html), 0644)
			checkForErrors(err)
		}
	}

}

func readConfig(cfile string) {
	yamlbytes, err := ioutil.ReadFile(cfile)
	checkForErrors(err)
}
func createSite() {
	site := postprocess.NewSite()
	err = site.Read(string(yamlbytes))
	checkForErrors(err)
}

// read all html files in dir and add the menu
func addPages(dir) {
	files, err := ioutil.ReadDir(dir)
	checkForErrors(err)
}
func checkForErrors(err) {
	if err != nil {
		log.Fatal(err)
	}
}

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

	// read config
	yamlbytes, err := ioutil.ReadFile(cfile)

	if err != nil {
		log.Fatal(err)
	}

	// create Site
	site := postprocess.NewSite()
	err = site.Read(string(yamlbytes))

	if err != nil {
		log.Fatal(err)
	}

	// read all html files in dir and add the menu
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".html" {
			// read html file with {{.Nav}}
			templateBytes, err := ioutil.ReadFile(dir + "/" + f.Name())

			if err != nil {
				log.Fatal(err)
			}

			// insert Nav menu
			html, err := site.PostProcess(f.Name(), string(templateBytes))

			if err != nil {
				log.Fatal(err)
			}

			// write the html file
			err = ioutil.WriteFile(dir+"/"+f.Name(), []byte(html), 0644)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

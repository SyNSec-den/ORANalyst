// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	htmltemplate "html/template"
	"os"
	texttemplate "text/template"
)

const yamlAppsTemplateMarkdown = "{{ printf \"#ONOS Helm Chart Releases\"}}\n\n" +
	"{{range $key, $value := .Entries }}" +
	"{{ printf \"## %s\" $key }}\n\n" +
	"{{range $value}}" +
	"{{printf \"#### Version **%s**\" .Version}}\n" +
	"{{printf \"> Generated %s\" .Created}}\n\n" +
	"{{printf \"App Version **%s**\" .AppVersion}}\n\n" +
	"{{range .Urls}}" +
	"{{printf \"[%s](%s)\" . .}}\n" +
	"{{end}}\n\n" +
	"{{end}}\n" +
	"{{end}}\n"

const yamlAppsTemplateHTML = "" +
	"<!DOCTYPE html \n" +
	"PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \n" +
	"\"DTD/xhtml1-strict.dtd\">\n" +
	"<html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"en\" lang=\"en\">\n" +
	"\t<head>\n" +
	"\t\t<title>ONOS helm Chart Releases</title>\n" +
	"\t</head>\n" +
	"\t<body>\n" +
	"\t\t<h1>{{ printf \"ONOS Helm Chart Releases\"}}</h1>\n" +
	"{{range $key, $value := .Entries }}" +
	"\t\t<h2>{{ printf \"%s\" $key }}</h2>\n" +
	"{{range $value}}" +
	"\t\t<div id=\"{{printf \"%s-%s\" $key .Version}}\">\n" +
	"\t\t\t<h3>Version{{printf \"%s\" .Version}}</h3>\n" +
	"\t\t\t<p>{{printf \"Generated %s\" .Created}}</p>\n" +
	"\t\t\t<p>App Version <b>{{printf \"%s\" .AppVersion}}</b></p>\n" +
	"{{range .Urls}}" +
	"\t\t\t<a href=\"{{printf \"%s\" .}}\">{{printf \"%s\" .}}</a>\n" +
	"{{end}}\n" +
	"\t\t</div>\n" +
	"{{end}}\n" +
	"{{end}}</body></html>"

// Chart :
type Chart struct {
	APIVersion  string   `yaml:"apiVersion"`
	AppVersion  string   `yaml:"appVersion"`
	Version     string   `yaml:"version"`
	Created     string   `yaml:"created"`
	Description string   `yaml:"description"`
	Urls        []string `yaml:"urls"`
}

// IndexYaml :
type IndexYaml struct {
	APIVersion string             `yaml:"apiVersion"`
	Entries    map[string][]Chart `yaml:"entries"`
	Generated  string             `yaml:"generated"`
}

/**
 * A simple application that takes the generated index.yaml and outputs it in
 * a Markdown format - usually we pipe this to README.md when in the gh-pages branch
 */
func main() {
	file := flag.String("file", "index", "name of YAML file to parse (without extension or path)")
	htmlout := flag.Bool("html", false, "output HTML instead of Markdown")
	flag.Parse()
	indexYaml, err := getIndexYaml(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load %s.yaml %s\n", *file, err)
		os.Exit(1)
	}

	if *htmlout {
		tmplAppsList, _ := htmltemplate.New("yamlAppsTemplateMarkdown").Parse(yamlAppsTemplateMarkdown)
		err = tmplAppsList.Execute(os.Stdout, indexYaml)
	} else {
		tmplAppsList, _ := texttemplate.New("yamlAppsTemplateHtml").Parse(yamlAppsTemplateHTML)
		err = tmplAppsList.Execute(os.Stdout, indexYaml)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute %v\n", err)
		os.Exit(1)
	}
}

func getIndexYaml(location string) (IndexYaml, error) {
	indexYaml := &IndexYaml{}
	viper.SetConfigName(location)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return IndexYaml{}, err
	}

	if err := viper.Unmarshal(indexYaml); err != nil {
		return IndexYaml{}, err
	}

	return *indexYaml, nil
}

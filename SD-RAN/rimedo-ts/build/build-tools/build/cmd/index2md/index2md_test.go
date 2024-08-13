// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"gotest.tools/assert"
	htmltemplate "html/template"
	"strings"
	"testing"
	texttemplate "text/template"
	//texttemplate "text/template"
)

func Test_convertYaml(t *testing.T) {
	index, err := getIndexYaml("sample-index")
	assert.NilError(t, err, "Unexpected error loading YAML")
	assert.Equal(t, 8, len(index.Entries))

	tmplAppsListText, _ := texttemplate.New("yamlAppsTemplateMarkdown").Parse(yamlAppsTemplateMarkdown)
	markdownBuffer := new(bytes.Buffer)
	err = tmplAppsListText.Execute(markdownBuffer, index)
	assert.NilError(t, err)
	assert.Equal(t, 10538, len(markdownBuffer.String()))
	assert.Assert(t, strings.HasPrefix(markdownBuffer.String(), "#ONOS Helm Chart Releases"))

	tmplAppsListHTML, _ := htmltemplate.New("yamlAppsTemplateMarkdown").Parse(yamlAppsTemplateHTML)
	xhtmlBuffer := new(bytes.Buffer)
	err = tmplAppsListHTML.Execute(xhtmlBuffer, index)
	assert.NilError(t, err)
	assert.Equal(t, 13374, len(xhtmlBuffer.String()))
	assert.Assert(t, strings.HasPrefix(xhtmlBuffer.String(), "<!DOCTYPE html"))

}

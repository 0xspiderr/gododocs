package main

import (
	"encoding/xml"
	"log"
	"os"
	"strings"
	"text/template"
)

type Class struct {
	XMLName          xml.Name   `xml:"class"`
	Name             string     `xml:"name,attr"`
	Inherits         string     `xml:"inherits,attr"`
	BriefDescription string     `xml:"brief_description"`
	Description      string     `xml:"description"`
	Tutorials        string     `xml:"tutorials"`
	Methods          []*Method  `xml:"methods>method"`
	Members          []Member   `xml:"members>member"`
	Constants        []Constant `xml:"constants>constant"`
	Signals          []Signal   `xml:"signals>signal"`
}

type Method struct {
	XMLName     xml.Name `xml:"method"`
	Name        string   `xml:"name,attr"`
	Qualifiers  string   `xml:"qualifiers,attr"`
	Return      Return   `xml:"return"`
	Param       []Param  `xml:"param"`
	Description string   `xml:"description"`
}

type Return struct {
	XMLName xml.Name `xml:"return"`
	Type    string   `xml:"type,attr"`
	Enum    string   `xml:"enum,attr"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Index   string   `xml:"index,attr"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

type Member struct {
	XMLName     xml.Name `xml:"member"`
	Description string   `xml:",chardata"`
	Name        string   `xml:"name,attr"`
	Type        string   `xml:"type,attr"`
	Setter      string   `xml:"setter,attr"`
	Getter      string   `xml:"getter,attr"`
}

type Constant struct {
	XMLName xml.Name `xml:"constant"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
	Enum    string   `xml:"enum,attr"`
}

type Signal struct {
	XMLName     xml.Name `xml:"signal"`
	Name        string   `xml:"name,attr"`
	Params      []Param  `xml:"param"`
	Description string   `xml:"description"`
}

// purpose: Joins parameters by name and type and then separates
// them by comma. This method makes templates lighter. It is
// necessary also because both Method[] and Signal[] have []Param
// and I didnt want to repeat the code for both of them.
func joinParameters(params []Param) string {
	pr := make([]string, len(params))
	for i, p := range params {
		pr[i] += p.Name + ": " + p.Type
	}
	return strings.Join(pr, ", ")
}

// purpose: Parses template files and structures them to
// output a markdown file with the documentation of the class.
func parseTemplate(c Class) {
	fm := template.FuncMap{"joinParameters": joinParameters}
	t := template.New("main.tmpl").Funcs(fm)
	t = template.Must(t.ParseFiles("main.tmpl", "class.tmpl", "members.tmpl", "methods.tmpl", "signals.tmpl"))

	f, err := os.Create("test.md")
	if err != nil {
		panic(err)
	}
	t.Execute(f, c)
}

// todo:
// 1. add flag commands
// -dir input directory with .xml files
// -out output directory where .md files will be stored
// 2. crawl -dir and subdirs
// 3. optional format [-format] .md/.html
// 4. maybe change template formatting to look different
func main() {
	f, err := os.ReadFile("NetCharacterComponent.xml")
	if err != nil {
		log.Fatal(err)
	}
	var c Class
	err = xml.Unmarshal(f, &c)
	if err != nil {
		log.Fatal(err)
	}
	parseTemplate(c)
}

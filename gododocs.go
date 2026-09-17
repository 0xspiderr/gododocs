package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// PLAN
// 1. read input xml file(s) from a given directory sequentially
// 2. convert xml files of godot documentation to markdown
// 2.1 -> implies I have to map the generate xml to a markdown structure
// 2.2 -> optionally expand the tool to convert to other formats
// 3. output to a given directory from cmdline

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
	ParamString string
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

////////////////////////////////////////////////////////////////////////
// func parseXML(c Class) string {									  //
// 	var class string												  //
// 	class = fmt.Sprintf("# %s Class Reference\n"+					  //
// 		"Inherits %s" +												  //
// 		"## Synopsis\n"+											  //
// 		"```gdscript "+												  //
// 		"class_name %s```"+											  //
// 		"%s\n", c.Name, c.Inherits, c.BriefDescription, c.Name)		  //
// 																	  //
// 	var members string												  //
// 	for _, m := range c.Members {									  //
// 		members += fmt.Sprintf("## Members\n" +						  //
// 			"```gdscript " +										  //
// 			"%s:%s```\n" , m.Name, m.Type) 							  //
// 	}																  //
// 																	  //
// }																  //
////////////////////////////////////////////////////////////////////////

// purpose: Joins parameters by name and type and then separates
// them by comma. This method makes templates lighter.
func (c *Class) aggregateParameters() {
	for _, m := range c.Methods {
		params := make([]string, len(m.Param))
		for i, p := range m.Param {
			params[i] += p.Name + ": " + p.Type
		}
		m.ParamString = strings.Join(params, ", ")
	}
}

// purpose: Parses template files and structures them to
// output a markdown file.
func parseTemplate(c Class) {
	c.aggregateParameters()
	t, err := template.ParseFiles("methods.tmpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("test.md") //

	if err != nil {
		panic(err)
	}
	err = t.Execute(f, c)
}

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
	fmt.Println()
	parseTemplate(c)
}

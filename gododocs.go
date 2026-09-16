package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
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
	Methods          []Method   `xml:"methods>method"`
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
	XMLName     xml.Name `xml:"constant"`
	Name        string   `xml:"name,attr"`
	Params      []Param  `xml:"param"`
	Description string   `xml:"description"`
}

func main() {
	f, err := os.ReadFile("./NetCharacterComponent.xml")
	if err != nil {
		log.Fatal(err)
	}
	var d Class
	err = xml.Unmarshal(f, &d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
}

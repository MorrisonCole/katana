package main

import "encoding/xml"

type JMDict struct {
	XMLName xml.Name `xml:"JMdict"`
	Entry []Entry `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name `xml:"entry"`
}
package model

import "encoding/xml"

type JMDict struct {
	XMLName xml.Name `model:"JMdict"`
	Entry   []Entry  `model:"entry"`
}

type Entry struct {
	XMLName xml.Name `model:"entry"`
}

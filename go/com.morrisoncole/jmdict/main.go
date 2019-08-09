package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

const jmDictPath = "jmdict.xml"

func main() {

	jmDictXml, err := os.Open(jmDictPath)

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jmDictXml)

	var jmDict JMDict
	xmlErr := xml.Unmarshal(byteValue, &jmDict)

	if xmlErr != nil {
		log.Fatal(xmlErr)
	}

	defer jmDictXml.Close()
}

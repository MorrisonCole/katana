package ingest

import (
	"bytes"
	"com.morrisoncole/katana/ingest/model"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const jmDictPath = "jmdict.xml"

func Ingest() {
	jmDictXml, err := os.Open(jmDictPath)

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jmDictXml)
	var jmDict model.JMDict

	d := xml.NewDecoder(bytes.NewReader(byteValue))

	d.Entity = model.EntityMap
	xmlErr := d.Decode(&jmDict)
	if xmlErr != nil {
		fmt.Printf("error: %v", xmlErr)
		return
	}

	out, err := json.MarshalIndent(jmDict, "", "  ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	ioutil.WriteFile("jmdict.json", out, 0644)
}

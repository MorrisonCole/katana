package ingest

import (
	"bytes"
	"com.morrisoncole/katana/ingest/model"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

func Ingest(path string) *model.JMDict {
	jmDictXml, err := os.Open(path + ".xml")

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jmDictXml)
	var jmDict model.JMDict

	d := xml.NewDecoder(bytes.NewReader(byteValue))

	d.Entity = model.EntityMap
	xmlErr := d.Decode(&jmDict)
	if xmlErr != nil {
		log.Fatalf("error: %v", xmlErr)
	}

	out, err := json.MarshalIndent(jmDict, "", "  ")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	ioutil.WriteFile(path+"_output.json", out, 0644)

	return &jmDict
}

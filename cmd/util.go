package templater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

// Replace function
func Replace() {

	if len(os.Args) != 3 {
		log.Fatal(fmt.Errorf("This util can only accept exactly 2 parameters. \n \t1. Template File \n \t2. Values File - must always be in valid json format \n Sample: \n \t templater template.md values.json"))
	}

	templateFile := os.Args[1]
	valuesFile := os.Args[2]

	jsonFile, err := os.Open(valuesFile)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file \"%s\". %s", valuesFile, err))
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonDef map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &jsonDef)
	if err != nil {
		log.Fatal(fmt.Errorf("Error unmarshalling json file \"%s\". %s", valuesFile, err))
	}

	err = jsonFile.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("Error closing json file \"%s\". %s", valuesFile, err))
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(fmt.Errorf("Error parsing template file \"%s\". %s", templateFile, err))
	}

	err = tmpl.ExecuteTemplate(os.Stdout, templateFile, jsonDef)
	if err != nil {
		log.Fatal(fmt.Errorf("Error executing template file \"%s\". %s", templateFile, err))
	}

}

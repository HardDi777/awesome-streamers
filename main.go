package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type LiveStream struct {
	Category  string     `json:"category" yaml:"category"`
	Streamers []Streamer `json:"streamers" yaml:"streamers"`
}

type Streamer struct {
	Name    string   `json:"name" yaml:"name"`
	Social  string   `json:"social" yaml:"social"`
	Streams Streams  `json:"streams" yaml:"streams"`
	Topics  []string `json:"topics" yaml:"topics"`
}
	defer fin.Close()

	data, err := ioutil.ReadAll(fin)
	if err != nil {
		log.Fatalln("cannot read file:", err)
	}

	var ls []LiveStream
	err = yaml.Unmarshal(data, &ls)
	if err != nil {
		log.Fatalln("cannot unmarshal YAML:", err)
	}

	json, err := json.MarshalIndent(ls, "", "  ")
	if err != nil {
		log.Fatalln("cannout marshal JSON:", err)
	}

	err = ioutil.WriteFile("awesome-streamers.json", json, 0644)
	if err != nil {
		log.Fatalln("cannot write JSON:", err)
	}

	tmpl, err := template.New("readme.tmpl").ParseFiles("template/readme.tmpl")
	if err != nil {
		log.Fatalln("cannot parse template", err)
	}

	fout, err := os.Create("README.md")
	if err != nil {
		log.Fatalln("cannot create README:", err)
	}
	defer fout.Close()

	w := bufio.NewWriter(fout)
	err = tmpl.Execute(w, ls)
	if err != nil {
		log.Fatalln("cannot execute template:", err)
	}
	defer w.Flush()
}
